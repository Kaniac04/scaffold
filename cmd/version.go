package cmd

import (
	"fmt"
	"os"
	"scaffold/internal"
)

func DisplayVer() {
	fmt.Fprintf(os.Stdout, "%sCurrent Scaffold Version : %s%s\n", internal.BlueCode, internal.Version, internal.ResetCode)
}
