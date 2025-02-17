package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"scaffold/internal"
	"scaffold/templates"
)

func Climb(ProjName string, FlagVal []string, CurrentWD string) {
	ProjType := flag.String("type", "", "project type")
	flag.CommandLine.Parse(FlagVal)

	fmt.Fprintf(os.Stdout, "Parsed Arguments...\n\t\tProject Name : %s%s%s\n\t\tProject Type : %s%s%s\n", internal.BlueCode, ProjName, internal.ResetCode, internal.BlueCode, *ProjType, internal.ResetCode)

	if _, exists := templates.AvailableProjects[*ProjType]; !exists {
		internal.PrintInvalidType(*ProjType)
		os.Exit(1)
	}

	ChosenDirectory := templates.ProjectDirectories[*ProjType]

	for _, direc_name := range ChosenDirectory {
		if err := os.MkdirAll(path.Join(CurrentWD, ProjName, direc_name), os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%s[ERROR] Error in creating directory : %s%s\n", internal.RedCode, direc_name, internal.ResetCode)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stdout, "%sScaffold Construction Complete!%s\n", internal.GreenCode, internal.ResetCode)
}
