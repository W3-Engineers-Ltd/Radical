package main

import (
	"flag"
	"log"
	"os"

	"github.com/W3-Engineers-Ltd/Radical/cmd"
	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	"github.com/W3-Engineers-Ltd/Radical/config"
	"github.com/W3-Engineers-Ltd/Radical/utils"
)

func main() {
	utils.NoticeUpdateradical()
	flag.Usage = cmd.Usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()

	if len(args) < 1 {
		cmd.Usage()
		os.Exit(2)
		return
	}

	if args[0] == "help" {
		cmd.Help(args[1:])
		return
	}

	for _, c := range commands.AvailableCommands {
		if c.Name() == args[0] && c.Run != nil {
			c.Flag.Usage = func() { c.Usage() }
			if c.CustomFlags {
				args = args[1:]
			} else {
				c.Flag.Parse(args[1:])
				args = c.Flag.Args()
			}

			if c.PreRun != nil {
				c.PreRun(c, args)
			}

			config.LoadConfig()
			os.Exit(c.Run(c, args))
			return
		}
	}

	utils.PrintErrorAndExit("Unknown subcommand", cmd.ErrorTemplate)
}
