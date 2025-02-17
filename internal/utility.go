package internal

import "fmt"

const Version = "v1.0"

func PrintGenericError() {
	fmt.Printf("[ERROR] Incorrect usage.\nYou might wanna refer help. Use : scaffold help\n")

}

func PrintMissingArgsError(command string) {
	fmt.Printf("[ERROR] Incorrect usage. Missing arguments for command: %s.\nRefer help: scaffold help\n", command)

}

func PrintBadArgsError(command string) {
	fmt.Printf("[ERROR] Incorrect usage. Blank Argument or Missing flag for command: %s.\nRefer help: scaffold help\n", command)
}

func PrintInvalidType(ProjType string) {
	fmt.Printf("[ERROR] Invalid Project Type.\nProject Type [%s] does not exist.\nRefer project list: scaffold list\n", ProjType)

}
