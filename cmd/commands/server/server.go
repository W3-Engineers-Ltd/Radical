package apiapp

import (
	"net/http"

	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"

	"os"

	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	"github.com/W3-Engineers-Ltd/Radical/cmd/commands/version"
	"github.com/W3-Engineers-Ltd/Radical/utils"
)

var CmdServer = &commands.Command{
	// CustomFlags: true,
	UsageLine: "server [port]",
	Short:     "serving static content over HTTP on port",
	Long: `
  The command 'server' creates a radiant API application.
`,
	PreRun: func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() },
	Run:    createAPI,
}

var (
	a utils.DocValue
	p utils.DocValue
	f utils.DocValue
)

func init() {
	CmdServer.Flag.Var(&a, "a", "Listen address")
	CmdServer.Flag.Var(&p, "p", "Listen port")
	CmdServer.Flag.Var(&f, "f", "Static files fold")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdServer)
}

func createAPI(cmd *commands.Command, args []string) int {
	if len(args) > 0 {
		err := cmd.Flag.Parse(args[1:])
		if err != nil {
			radicalLogger.Log.Error(err.Error())
		}
	}
	if a == "" {
		a = "127.0.0.1"
	}
	if p == "" {
		p = "8080"
	}
	if f == "" {
		cwd, _ := os.Getwd()
		f = utils.DocValue(cwd)
	}
	radicalLogger.Log.Infof("Start server on http://%s:%s, static file %s", a, p, f)
	err := http.ListenAndServe(string(a)+":"+string(p), http.FileServer(http.Dir(f)))
	if err != nil {
		radicalLogger.Log.Error(err.Error())
	}
	return 0
}
