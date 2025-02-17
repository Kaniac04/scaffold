package cmd

import (
	"fmt"
	"os"
	"github.com/Kaniac04/scaffold/internal"
)

func DisplayVer() {
	fmt.Fprintf(os.Stdout, "%sCurrent Scaffold Version : %s%s\n", internal.BlueCode, internal.Version, internal.ResetCode)
}
