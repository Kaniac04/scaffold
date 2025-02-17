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
			fmt.Fprintf(os.Stdout, "%sAvailable Project Templates : %s\n\n", internal.CyanCode, internal.ResetCode)
			for key, value := range templates.ProjectTemplates {
				fmt.Fprintf(os.Stdout, "%sProject Type : %s\n%s%s\n\n", internal.CyanCode, key, value, internal.ResetCode)
			}
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stdout, "%sProject Type : %s\n%s%s", internal.CyanCode, *ProjType, templates.ProjectTemplates[*ProjType], internal.ResetCode)
			os.Exit(0)
		}
	}
}
