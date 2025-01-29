package main

import (
	"fmt"
	"log"
	"os"

	"kcommit/src"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Println(src.KcVersion)
			return
		}
	}

	fileManager, err := src.NewFileManager()
	if err != nil {
		log.Fatalln(err, "Failed to initialize FileManager")
	}

	git := src.NewGit()
	utils := src.NewUtils()
	viewBuilder := src.NewViewBuilder()

	runner := src.NewRunner(fileManager, git, utils, viewBuilder)

	runner.Start()
}
