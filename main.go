package main

import (
	"os"

	"github.com/Kaniac04/scaffold/cmd"
	"github.com/Kaniac04/scaffold/internal"
)

var CommandList = map[string]struct{}{
	"help":    {},
	"climb":   {},
	"list":    {},
	"version": {},
	"view":    {},
}

func main() {

	if _, exists := CommandList[os.Args[1]]; !exists {
		internal.PrintGenericError()
		os.Exit(1)
	}

	CWD, _ := os.Getwd()

	CommandString := os.Args[1:]
	command := CommandString[0]

	switch command {
	case "help":
		if len(CommandString) > 1 {
			internal.PrintGenericError()
			os.Exit(1)
		}
		cmd.PrintGlobalHelp()
		os.Exit(0)

	case "climb":
		if len(CommandString) != 4 {
			internal.PrintMissingArgsError(command)
			os.Exit(1)
		} else if CommandString[2] != "--type" && CommandString[2] != "-type" || CommandString[1] == " " || CommandString[1] == "" {
			internal.PrintBadArgsError(command)
			os.Exit(1)
		} else {
			cmd.Climb(CommandString[1], CommandString[2:], CWD)
			os.Exit(0)
		}

	case "version":
		if len(CommandString) > 1 {
			internal.PrintGenericError()
			os.Exit(1)
		} else {
			cmd.DisplayVer()
			os.Exit(0)
		}

	case "list":
		if len(CommandString) > 1 {
			internal.PrintGenericError()
			os.Exit(1)
		} else {
			cmd.ListTemp()
			os.Exit(0)
		}

	case "view":
		if len(CommandString) == 1 {
			cmd.ViewStruct([]string{"--type", "all"})
			os.Exit(0)
		} else {
			cmd.ViewStruct(CommandString[1:])
			os.Exit(0)
		}
	}

}
