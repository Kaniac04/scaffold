package cmd

import (
	"os"

	"github.com/Kaniac04/scaffold/internal"
	"github.com/fatih/color"
)

func DisplayVer() {
	color.New(color.FgBlue).Fprintf(os.Stdout, "Current Scaffold Version : %s", internal.Version)
	color.New(color.Reset).Fprintf(os.Stdout, "%s\n", "")
}
