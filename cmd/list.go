package cmd

import (
	"fmt"

	"github.com/Kaniac04/scaffold/templates"
)

func ListTemp() {

	fmt.Printf("Available Projects : \n\n")

	for ProjName := range templates.AvailableProjects {
		fmt.Printf("%s\n", ProjName)

	}
}
