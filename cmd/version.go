package cmd

import (
	"fmt"
	"os"

	"github.com/Kaniac04/scaffold/internal"
)

func DisplayVer() {
	fmt.Fprintf(os.Stdout, "Current Scaffold Version : %s", internal.Version)

}
