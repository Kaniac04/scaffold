package cmd

import (
	"fmt"
	"os"

	"github.com/Kaniac04/scaffold/templates"
)

func ListTemp() {

	fmt.Fprintf(os.Stdout, "Available Projects : \n")

	for ProjName := range templates.AvailableProjects {
		fmt.Fprintf(os.Stdout, "%s\n", ProjName)

	}
}
