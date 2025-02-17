package internal

import (
	"fmt"
	"os"
)

const Version = "v1.0"

const (
	ResetCode  = "\033[0m"
	RedCode    = "\033[31m"
	BlueCode   = "\033[34m"
	GreenCode  = "\033[32m"
	YellowCode = "\033[33m"
	CyanCode   = "\033[36m"
)

func PrintGenericError() {
	fmt.Fprintf(os.Stderr, "%s[ERROR] Incorrect usage.%s\nYou might wanna refer help. Use : scaffold help%s\n", RedCode, YellowCode, ResetCode)
}

func PrintMissingArgsError(command string) {
	fmt.Fprintf(os.Stderr, "%s[ERROR] Incorrect usage.%s\nMissing arguments for command : %s.\nRefer help : scaffold help%s\n", RedCode, YellowCode, command, ResetCode)
}

func PrintBadArgsError(command string) {
	fmt.Fprintf(os.Stderr, "%s[ERROR] Incorrect usage.%s\nBlank Argument or Missing flag for command : %s.\nRefer help : scaffold help%s\n", RedCode, YellowCode, command, ResetCode)
}

func PrintInvalidType(ProjType string) {
	fmt.Fprintf(os.Stderr, "%s[ERROR] Invalid Project Type.\nProject Type [%s] does not exists%s.\nRefer project list : scaffold list%s\n", RedCode, ProjType, YellowCode, ResetCode)
}
