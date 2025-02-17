package cmd

import (
	"fmt"

	"github.com/Kaniac04/scaffold/internal"
)

func DisplayVer() {
	fmt.Printf("Current Scaffold Version : %s\n", internal.Version)

}
