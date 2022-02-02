package version

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	path "path/filepath"
	"regexp"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	"github.com/W3-Engineers-Ltd/Radical/config"
	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
	"github.com/W3-Engineers-Ltd/Radical/logger/colors"
	"github.com/W3-Engineers-Ltd/Radical/utils"
)

const verboseVersionBanner string = `%s%s______
██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v{{ .radicalVersion }}%s
%s%s
├── radiant     : {{ .radiantVersion }}
├── GoVersion : {{ .GoVersion }}
├── GOOS      : {{ .GOOS }}
├── GOARCH    : {{ .GOARCH }}
├── NumCPU    : {{ .NumCPU }}
├── GOPATH    : {{ .GOPATH }}
├── GOROOT    : {{ .GOROOT }}
├── Compiler  : {{ .Compiler }}
└── Date      : {{ Now "Monday, 2 Jan 2006" }}%s
`

const shortVersionBanner = `______
██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v{{ .radicalVersion }}
`

var CmdVersion = &commands.Command{
	UsageLine: "version",
	Short:     "Prints the current radical version",
	Long: `
Prints the current radical, radiant and Go version alongside the platform information.
`,
	Run: versionCmd,
}
var outputFormat string

const version = config.Version

func init() {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	fs.StringVar(&outputFormat, "o", "", "Set the output format. Either json or yaml.")
	CmdVersion.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, CmdVersion)
}

func versionCmd(cmd *commands.Command, args []string) int {

	cmd.Flag.Parse(args)
	stdout := cmd.Out()

	if outputFormat != "" {
		runtimeInfo := RuntimeInfo{
			GetGoVersion(),
			runtime.GOOS,
			runtime.GOARCH,
			runtime.NumCPU(),
			os.Getenv("GOPATH"),
			runtime.GOROOT(),
			runtime.Compiler,
			version,
			GetradiantVersion(),
		}
		switch outputFormat {
		case "json":
			{
				b, err := json.MarshalIndent(runtimeInfo, "", "    ")
				if err != nil {
					radicalLogger.Log.Error(err.Error())
				}
				fmt.Println(string(b))
				return 0
			}
		case "yaml":
			{
				b, err := yaml.Marshal(&runtimeInfo)
				if err != nil {
					radicalLogger.Log.Error(err.Error())
				}
				fmt.Println(string(b))
				return 0
			}
		}
	}

	coloredBanner := fmt.Sprintf(verboseVersionBanner, "\x1b[35m", "\x1b[1m",
		"\x1b[0m", "\x1b[32m", "\x1b[1m", "\x1b[0m")
	InitBanner(stdout, bytes.NewBufferString(coloredBanner))
	return 0
}

// ShowShortVersionBanner prints the short version banner.
func ShowShortVersionBanner() {
	output := colors.NewColorWriter(os.Stdout)
	InitBanner(output, bytes.NewBufferString(colors.MagentaBold(shortVersionBanner)))
}

func GetradiantVersion() string {
	re, err := regexp.Compile(`VERSION = "([0-9.]+)"`)
	if err != nil {
		return ""
	}
	wgopath := utils.GetGOPATHs()
	if len(wgopath) == 0 {
		radicalLogger.Log.Error("GOPATH environment is empty,may be you use `go module`")
		return ""
	}
	for _, wg := range wgopath {
		wg, _ = path.EvalSymlinks(path.Join(wg, "src", "github.com", "astaxie", "radiant"))
		filename := path.Join(wg, "radiant.go")
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			radicalLogger.Log.Error("Error while getting stats of 'radiant.go'")
		}
		fd, err := os.Open(filename)
		if err != nil {
			radicalLogger.Log.Error("Error while reading 'radiant.go'")
			continue
		}
		reader := bufio.NewReader(fd)
		for {
			byteLine, _, er := reader.ReadLine()
			if er != nil && er != io.EOF {
				return ""
			}
			if er == io.EOF {
				break
			}
			line := string(byteLine)
			s := re.FindStringSubmatch(line)
			if len(s) >= 2 {
				return s[1]
			}
		}

	}
	return "radiant is not installed. Please do consider installing it first: https://github.com/W3-Engineers-Ltd/Radiant. " +
		"If you are using go mod, and you don't install the radiant under $GOPATH/src/github.com/astaxie, just ignore this."
}

func GetGoVersion() string {
	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command("go", "version").Output(); err != nil {
		radicalLogger.Log.Fatalf("There was an error running 'go version' command: %s", err)
	}
	return strings.Split(string(cmdOut), " ")[2]
}
