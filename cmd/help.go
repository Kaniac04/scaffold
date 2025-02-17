package cmd

import (
	"fmt"
	"os"
)

func PrintGlobalHelp() {
	fmt.Fprintf(os.Stdout, `

░██████╗░█████╗░░█████╗░███████╗███████╗░█████╗░██╗░░░░░██████╗░
██╔════╝██╔══██╗██╔══██╗██╔════╝██╔════╝██╔══██╗██║░░░░░██╔══██╗
╚█████╗░██║░░╚═╝███████║█████╗░░█████╗░░██║░░██║██║░░░░░██║░░██║
░╚═══██╗██║░░██╗██╔══██║██╔══╝░░██╔══╝░░██║░░██║██║░░░░░██║░░██║
██████╔╝╚█████╔╝██║░░██║██║░░░░░██║░░░░░╚█████╔╝███████╗██████╔╝
╚═════╝░░╚════╝░╚═╝░░╚═╝╚═╝░░░░░╚═╝░░░░░░╚════╝░╚══════╝╚═════╝░

A lightweight and efficient CLI tool to generate organized directory structures for various project types.  
This tool helps developers quickly scaffold industry-evolved project structures for Flask, Gin, etc.

Usage:  
  scaffold [command] [flags]  

Commands:  
  climb     Initialize a new project with the desired structure in the current directory  
  list      Show available project types
  view      Show the directory structure for a project type
  version   Display the CLI version  
  help      Show this help message  

Usage (climb):
  scaffold climb <project_name> --type <project_type>

Usage (list):
  scaffold list

Usage (view):
  scaffold view --type <project_type>

Usage (version):
  scaffold version

Usage (help):
  scaffold help
`)
}
