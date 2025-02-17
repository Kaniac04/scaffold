package cmd

import (
	"flag"
	"os"
	"path"

	"github.com/Kaniac04/scaffold/internal"
	"github.com/Kaniac04/scaffold/templates"
	"github.com/fatih/color"
)

func Climb(ProjName string, FlagVal []string, CurrentWD string) {
	ProjType := flag.String("type", "", "project type")
	flag.CommandLine.Parse(FlagVal)

	color.New(color.FgBlue).Fprintf(os.Stdout, "Parsed Arguments...\n")
	color.New(color.FgBlue).Fprintf(os.Stdout, "\t\tProject Name : %s", ProjName)
	color.New(color.Reset).Fprintf(os.Stdout, "%s\n", "")
	color.New(color.FgBlue).Fprintf(os.Stdout, "\t\tProject Type : %s", *ProjType)
	color.New(color.Reset).Fprintf(os.Stdout, "%s\n", "")

	if _, exists := templates.AvailableProjects[*ProjType]; !exists {
		internal.PrintInvalidType(*ProjType)
		os.Exit(1)
	}

	ChosenDirectory := templates.ProjectDirectories[*ProjType]

	for _, direc_name := range ChosenDirectory {
		if err := os.MkdirAll(path.Join(CurrentWD, ProjName, direc_name), os.ModePerm); err != nil {
			color.New(color.FgRed).Fprintf(os.Stderr, "[ERROR] Error in creating directory : %s", direc_name)
			color.New(color.Reset).Fprintf(os.Stderr, "%s\n", "")
			os.Exit(1)
		}
	}

	color.New(color.FgGreen).Fprintf(os.Stdout, "Scaffold Construction Complete!")
	color.New(color.Reset).Fprintf(os.Stdout, "\n")
}
