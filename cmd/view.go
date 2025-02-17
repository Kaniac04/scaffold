package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/Kaniac04/scaffold/internal"
	"github.com/Kaniac04/scaffold/templates"
)

func ViewStruct(TypeFlag []string) {
	ProjType := flag.String("type", "all", "project type")
	flag.CommandLine.Parse(TypeFlag)

	if _, exists := templates.ProjectTemplates[*ProjType]; !exists && *ProjType != "all" {
		internal.PrintInvalidType(*ProjType)
		os.Exit(1)
	} else {
		if *ProjType == "all" {
			fmt.Fprintf(os.Stdout, "Available Project Templates :\n\n")

			for key, value := range templates.ProjectTemplates {
				fmt.Fprintf(os.Stdout, "Project Type : %s\n%s\n", key, value)
			}
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stdout, "Project Type : %s\n%s\n", *ProjType, templates.ProjectTemplates[*ProjType])

			os.Exit(0)
		}
	}
}
