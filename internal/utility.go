package internal

import "fmt"

const Version = "v1.0"

func PrintGenericError() {
	fmt.Printf(`[ERROR] Incorrect usage.
You might wanna refer help. Use : scaffold help
`)

}

func PrintMissingArgsError(command string) {
	fmt.Printf(`[ERROR] Incorrect usage. Missing arguments for command: %s.
Refer help: scaffold help
`, command)

}

func PrintBadArgsError(command string) {
	fmt.Printf(`[ERROR] Incorrect usage. Blank Argument or Missing flag for command: %s.
Refer help: scaffold help
`, command)
}

func PrintInvalidType(ProjType string) {
	fmt.Printf(`[ERROR] Invalid Project Type.
Project Type [%s] does not exist.
Refer project list: scaffold list
`, ProjType)

}
