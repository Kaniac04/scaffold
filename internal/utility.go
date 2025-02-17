package internal

import (
	"os"
	"github.com/fatih/color"
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
	color.New(color.FgRed).Fprintf(os.Stderr, "[ERROR] Incorrect usage.")
	color.New(color.FgYellow).Fprintf(os.Stderr, "\nYou might wanna refer help. Use : scaffold help")
	color.New(color.Reset).Fprintf(os.Stderr, "\n")
}

func PrintMissingArgsError(command string) {
	color.New(color.FgRed).Fprintf(os.Stderr, "[ERROR] Incorrect usage. ")
	color.New(color.FgYellow).Fprintf(os.Stderr, "Missing arguments for command: %s.\n", command)
	color.New(color.FgYellow).Fprintf(os.Stderr, "Refer help: scaffold help\n")
	color.New(color.Reset).Fprint(os.Stdout, "\n")
}

func PrintBadArgsError(command string) {
	color.New(color.FgRed).Fprintf(os.Stderr, "[ERROR] Incorrect usage. ")
	color.New(color.FgYellow).Fprintf(os.Stderr, "Blank Argument or Missing flag for command: %s.\n", command)
	color.New(color.FgYellow).Fprintf(os.Stderr, "Refer help: scaffold help\n")
	color.New(color.Reset).Fprint(os.Stdout, "\n")
}

func PrintInvalidType(ProjType string) {
	color.New(color.FgRed).Fprintf(os.Stderr, "[ERROR] Invalid Project Type.\n")
	color.New(color.FgYellow).Fprintf(os.Stderr, "Project Type [%s] does not exist.\n", ProjType)
	color.New(color.FgYellow).Fprintf(os.Stderr, "Refer project list: scaffold list\n")
	color.New(color.Reset).Fprint(os.Stdout, "\n")
}
