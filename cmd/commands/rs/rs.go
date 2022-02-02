// Package rs ...
package rs

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"strings"

	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	"github.com/W3-Engineers-Ltd/Radical/cmd/commands/version"
	"github.com/W3-Engineers-Ltd/Radical/config"
	"github.com/W3-Engineers-Ltd/Radical/logger"
	"github.com/W3-Engineers-Ltd/Radical/logger/colors"
	"github.com/W3-Engineers-Ltd/Radical/utils"
)

var cmdRs = &commands.Command{
	UsageLine: "rs",
	Short:     "Run customized scripts",
	Long: `Run script allows you to run arbitrary commands using radical.
  Custom commands are provided from the "scripts" object inside radical.json or radicalfile.

  To run a custom command, use: {{"$ radical rs mycmd ARGS" | bold}}
  {{if len .}}
{{"AVAILABLE SCRIPTS"|headline}}{{range $cmdName, $cmd := .}}
  {{$cmdName | bold}}
      {{$cmd}}{{end}}{{end}}
`,
	PreRun: func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() },
	Run:    runScript,
}

func init() {
	config.LoadConfig()
	cmdRs.Long = utils.TmplToString(cmdRs.Long, config.Conf.Scripts)
	commands.AvailableCommands = append(commands.AvailableCommands, cmdRs)
}

func runScript(cmd *commands.Command, args []string) int {
	if len(args) == 0 {
		cmd.Usage()
	}

	start := time.Now()
	script, args := args[0], args[1:]

	if c, exist := config.Conf.Scripts[script]; exist {
		command := customCommand{
			Name:    script,
			Command: c,
			Args:    args,
		}
		if err := command.run(); err != nil {
			radicalLogger.Log.Error(err.Error())
		}
	} else {
		radicalLogger.Log.Errorf("Command '%s' not found in radicalfile/radical.json", script)
	}
	elapsed := time.Since(start)
	fmt.Println(colors.GreenBold(fmt.Sprintf("Finished in %s.", elapsed)))
	return 0
}

type customCommand struct {
	Name    string
	Command string
	Args    []string
}

func (c *customCommand) run() error {
	radicalLogger.Log.Info(colors.GreenBold(fmt.Sprintf("Running '%s'...", c.Name)))
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin", "linux":
		args := append([]string{c.Command}, c.Args...)
		cmd = exec.Command("sh", "-c", strings.Join(args, " "))
	case "windows":
		args := append([]string{c.Command}, c.Args...)
		cmd = exec.Command("cmd", "/C", strings.Join(args, " "))
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
