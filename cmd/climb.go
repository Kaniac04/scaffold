package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Kaniac04/scaffold/internal"
	"github.com/Kaniac04/scaffold/templates"
)

func Climb(ProjName string, FlagVal []string, CurrentWD string) {
	ProjType := flag.String("type", "", "project type")
	flag.CommandLine.Parse(FlagVal)

	fmt.Fprintf(os.Stdout, "Parsed Arguments...\nProject Name : %s\nProject Type : %s\n", ProjName, *ProjType)

	if _, exists := templates.AvailableProjects[*ProjType]; !exists {
		internal.PrintInvalidType(*ProjType)
		os.Exit(1)
	}

	ChosenDirectory := templates.ProjectDirectories[*ProjType]

	for _, direc_path := range ChosenDirectory {
		path := path.Join(CurrentWD, ProjName, direc_path)
		if filepath.Ext(path) == "" {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", path, err)
				os.Exit(1)
			}
		} else {
			dir := filepath.Dir(path)
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
				os.Exit(1)
			}

			f, err := os.Create(path)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", path, err)
				os.Exit(1)
			} else {
				f.Close()
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Scaffold Construction Complete!")

}
