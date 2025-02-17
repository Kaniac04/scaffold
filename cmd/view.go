package cmd

import (
	"flag"
	"os"

	"github.com/Kaniac04/scaffold/internal"
	"github.com/Kaniac04/scaffold/templates"
	"github.com/fatih/color"
)

func ViewStruct(TypeFlag []string) {
	ProjType := flag.String("type", "all", "project type")
	flag.CommandLine.Parse(TypeFlag)

	if _, exists := templates.ProjectTemplates[*ProjType]; !exists && *ProjType != "all" {
		internal.PrintInvalidType(*ProjType)
		os.Exit(1)
	} else {
		if *ProjType == "all" {
			color.New(color.FgCyan).Fprintf(os.Stdout, "Available Project Templates : ")
			color.New(color.Reset).Fprintf(os.Stdout, "\n\n")
			for key, value := range templates.ProjectTemplates {
				color.New(color.FgCyan).Fprintf(os.Stdout, "Project Type : %s\n", key)
				color.New(color.Reset).Fprintf(os.Stdout, "%s\n\n", value)
			}
			os.Exit(0)
		} else {
			color.New(color.FgCyan).Fprintf(os.Stdout, "Project Type : %s\n", *ProjType)
			color.New(color.Reset).Fprintf(os.Stdout, "%s%s", templates.ProjectTemplates[*ProjType], "\n")
			os.Exit(0)
		}
	}
}
