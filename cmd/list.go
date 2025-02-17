package cmd

import (
	"os"

	"github.com/Kaniac04/scaffold/templates"
	"github.com/fatih/color"
)

func ListTemp() {

	color.New(color.FgBlue).Fprintf(os.Stdout, "Available Projects : \n\n")
	color.New(color.Reset).Fprintf(os.Stdout, "")
	for ProjName := range templates.AvailableProjects {
		color.New(color.FgBlue).Fprintf(os.Stdout, "%s", ProjName)
		color.New(color.Reset).Fprintf(os.Stdout, "")
	}
}
