package radiantpro

import (
	"strings"

	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	"github.com/W3-Engineers-Ltd/Radical/internal/app/module/radiantpro"
	"github.com/W3-Engineers-Ltd/Radical/logger"
)

var CmdradiantPro = &commands.Command{
	UsageLine: "pro [command]",
	Short:     "Source code generator",
	Long:      ``,
	Run:       radiantPro,
}

func init() {
	CmdradiantPro.Flag.Var(&radiantpro.SQL, "sql", "sql file path")
	CmdradiantPro.Flag.Var(&radiantpro.SQLMode, "sqlmode", "sql mode")
	CmdradiantPro.Flag.Var(&radiantpro.SQLModePath, "sqlpath", "sql mode path")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdradiantPro)
}

func radiantPro(cmd *commands.Command, args []string) int {
	if len(args) < 1 {
		radicalLogger.Log.Fatal("Command is missing")
	}

	if len(args) >= 2 {
		cmd.Flag.Parse(args[1:])
	}

	gcmd := args[0]
	switch gcmd {
	case "gen":
		radiantpro.DefaultradiantPro.Run()
	case "config":
		radiantpro.DefaultradiantPro.GenConfig()
	case "migration":
		radiantpro.DefaultradiantPro.Migration(args)
	default:
		radicalLogger.Log.Fatal("Command is missing")
	}
	radicalLogger.Log.Successf("%s successfully generated!", strings.Title(gcmd))
	return 0
}
