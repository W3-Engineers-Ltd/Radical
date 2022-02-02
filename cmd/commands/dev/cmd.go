// Copyright 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
