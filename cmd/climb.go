package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/Kaniac04/scaffold/internal"
	"github.com/Kaniac04/scaffold/templates"
)

func Climb(ProjName string, FlagVal []string, CurrentWD string) {
	ProjType := flag.String("type", "", "project type")
	flag.CommandLine.Parse(FlagVal)

	fmt.Fprintf(os.Stdout, "Parsed Arguments...\nProject Name : %s\nProject Type : %s\n", ProjName, *ProjType)

	if _, exists := templates.AvailableProjects[*ProjType]; !exists {
		internal.PrintInvalidType(*ProjType)
		os.Exit(1)
	}

	ChosenDirectory := templates.ProjectDirectories[*ProjType]

	for _, direc_name := range ChosenDirectory {
		if err := os.MkdirAll(path.Join(CurrentWD, ProjName, direc_name), os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Error in creating directory : %s", direc_name)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stderr, "Scaffold Construction Complete!")

}
