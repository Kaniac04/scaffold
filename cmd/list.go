package cmd

import (
	"fmt"
	"os"
	"github.com/Kaniac04/scaffold/internal"
	"github.com/Kaniac04/scaffold/templates"
)

func ListTemp() {

	fmt.Fprintf(os.Stdout, "%sAvailable Projects : \n\n%s", internal.BlueCode, internal.ResetCode)
	for ProjName := range templates.AvailableProjects {
		fmt.Fprintf(os.Stdout, "%s%s%s", internal.BlueCode, ProjName, internal.ResetCode)
	}
}
