package dev

import (
	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
)

var CmdDev = &commands.Command{
	CustomFlags: true,
	UsageLine:   "dev [command]",
	Short:       "Commands which used to help to develop radiant and radical",
	Long: `
Commands that help developer develop, build and test radiant.
- githook    Prepare githooks
`,
	Run: Run,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdDev)
}

func Run(cmd *commands.Command, args []string) int {
	if len(args) < 1 {
		radicalLogger.Log.Fatal("Command is missing")
	}

	if len(args) >= 2 {
		cmd.Flag.Parse(args[1:])
	}

	gcmd := args[0]

	switch gcmd {

	case "githook":
		initGitHook()
	default:
		radicalLogger.Log.Fatal("Unknown command")
	}
	return 0
}
