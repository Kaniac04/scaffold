package internal

import (
	"fmt"
	"os"
)

const Version = "v1.0"

func PrintGenericError() {
	fmt.Fprintf(os.Stdout, "[ERROR] Incorrect usage. You might wanna refer help. Use : scaffold help")

}

func PrintMissingArgsError(command string) {
	fmt.Fprintf(os.Stdout, "[ERROR] Incorrect usage. Missing arguments for command: %s. Refer help: scaffold help", command)

}

func PrintBadArgsError(command string) {
	fmt.Fprintf(os.Stdout, "[ERROR] Incorrect usage. Blank Argument or Missing flag for command: %s. Refer help: scaffold help", command)
}

func PrintInvalidType(ProjType string) {
	fmt.Fprintf(os.Stdout, "[ERROR] Invalid Project Type. Project Type [%s] does not exist. Refer project list: scaffold list", ProjType)

}
