// Package dlv ...
package dlv

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	"github.com/W3-Engineers-Ltd/Radical/cmd/commands/version"
	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
	"github.com/W3-Engineers-Ltd/Radical/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/go-delve/delve/pkg/terminal"
	"github.com/go-delve/delve/service"
	"github.com/go-delve/delve/service/debugger"
	"github.com/go-delve/delve/service/rpc2"
	"github.com/go-delve/delve/service/rpccommon"
)

var cmdDlv = &commands.Command{
	CustomFlags: true,
	UsageLine:   "dlv [-package=\"\"] [-port=8181] [-verbose=false]",
	Short:       "Start a debugging session using Delve",
	Long: `dlv command start a debugging session using debugging tool Delve.

  To debug your application using Delve, use: {{"$ radical dlv" | bold}}

  For more information on Delve: https://github.com/go-delve/delve
`,
	PreRun: func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() },
	Run:    runDlv,
}

var (
	packageName string
	verbose     bool
	port        int
)

func init() {
	fs := flag.NewFlagSet("dlv", flag.ContinueOnError)
	fs.StringVar(&packageName, "package", "", "The package to debug (Must have a main package)")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose mode")
	fs.IntVar(&port, "port", 8181, "Port to listen to for clients")
	cmdDlv.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, cmdDlv)
}

func runDlv(cmd *commands.Command, args []string) int {
	if err := cmd.Flag.Parse(args); err != nil {
		radicalLogger.Log.Fatalf("Error while parsing flags: %v", err.Error())
	}

	var (
		addr       = fmt.Sprintf("127.0.0.1:%d", port)
		paths      = make([]string, 0)
		notifyChan = make(chan int)
	)

	if err := loadPathsToWatch(&paths); err != nil {
		radicalLogger.Log.Fatalf("Error while loading paths to watch: %v", err.Error())
	}
	go startWatcher(paths, notifyChan)
	return startDelveDebugger(addr, notifyChan)
}

// buildDebug builds a debug binary in the current working directory
func buildDebug() (string, error) {
	args := []string{"-gcflags", "-N -l", "-o", "debug"}
	args = append(args, utils.SplitQuotedFields("-ldflags='-linkmode internal'")...)
	args = append(args, packageName)
	if err := utils.GoCommand("build", args...); err != nil {
		return "", err
	}

	fp, err := filepath.Abs("./debug")
	if err != nil {
		return "", err
	}
	return fp, nil
}

// loadPathsToWatch loads the paths that needs to be watched for changes
func loadPathsToWatch(paths *[]string) error {
	directory, err := os.Getwd()
	if err != nil {
		return err
	}
	filepath.Walk(directory, func(path string, info os.FileInfo, _ error) error {
		if strings.HasSuffix(info.Name(), "docs") {
			return filepath.SkipDir
		}
		if strings.HasSuffix(info.Name(), "swagger") {
			return filepath.SkipDir
		}
		if strings.HasSuffix(info.Name(), "vendor") {
			return filepath.SkipDir
		}

		if filepath.Ext(info.Name()) == ".go" {
			*paths = append(*paths, path)
		}
		return nil
	})
	return nil
}

// startDelveDebugger starts the Delve debugger server
func startDelveDebugger(addr string, ch chan int) int {
	radicalLogger.Log.Info("Starting Delve Debugger...")

	fp, err := buildDebug()
	if err != nil {
		radicalLogger.Log.Fatalf("Error while building debug binary: %v", err)
	}
	defer os.Remove(fp)

	abs, err := filepath.Abs("./debug")
	if err != nil {
		radicalLogger.Log.Fatalf("%v", err)
	}

	// Create and start the debugger server
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		radicalLogger.Log.Fatalf("Could not start listener: %s", err)
	}
	defer listener.Close()

	server := rpccommon.NewServer(&service.Config{
		Listener:    listener,
		AcceptMulti: true,
		APIVersion:  2,
		ProcessArgs: []string{abs},
		Debugger: debugger.Config{
			AttachPid:  0,
			WorkingDir: ".",
			Backend:    "default",
		},
	})
	if err := server.Run(); err != nil {
		radicalLogger.Log.Fatalf("Could not start debugger server: %v", err)
	}

	// Start the Delve client REPL
	client := rpc2.NewClient(addr)
	// Make sure the client is restarted when new changes are introduced
	go func() {
		for {
			if val := <-ch; val == 0 {
				if _, err := client.Restart(true); err != nil {
					utils.Notify("Error while restarting the client: "+err.Error(), "radical")
				} else {
					if verbose {
						utils.Notify("Delve Debugger Restarted", "radical")
					}
				}
			}
		}
	}()

	// Create the terminal and connect it to the client debugger
	term := terminal.New(client, nil)
	status, err := term.Run()
	if err != nil {
		radicalLogger.Log.Fatalf("Could not start Delve REPL: %v", err)
	}

	// Stop and kill the debugger server once user quits the REPL
	if err := server.Stop(); err != nil {
		radicalLogger.Log.Fatalf("Could not stop Delve server: %v", err)
	}
	return status
}

var eventsModTime = make(map[string]int64)

// startWatcher starts the fsnotify watcher on the passed paths
func startWatcher(paths []string, ch chan int) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		radicalLogger.Log.Fatalf("Could not start the watcher: %v", err)
	}
	defer watcher.Close()

	// Feed the paths to the watcher
	for _, path := range paths {
		if err := watcher.Add(path); err != nil {
			radicalLogger.Log.Fatalf("Could not set a watch on path: %v", err)
		}
	}

	for {
		select {
		case evt := <-watcher.Events:
			build := true
			if filepath.Ext(evt.Name) != ".go" {
				continue
			}

			mt := utils.GetFileModTime(evt.Name)
			if t := eventsModTime[evt.Name]; mt == t {
				build = false
			}
			eventsModTime[evt.Name] = mt

			if build {
				go func() {
					if verbose {
						utils.Notify("Rebuilding application with the new changes", "radical")
					}

					// Wait 1s before re-build until there is no file change
					scheduleTime := time.Now().Add(1 * time.Second)
					time.Sleep(time.Until(scheduleTime))
					_, err := buildDebug()
					if err != nil {
						utils.Notify("Build Failed: "+err.Error(), "radical")
					} else {
						ch <- 0 // Notify listeners
					}
				}()
			}
		case err := <-watcher.Errors:
			if err != nil {
				ch <- -1
			}
		}
	}
}
