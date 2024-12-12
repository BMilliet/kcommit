package main

import (
	"kcommit/src"
	"log"
)

func main() {
	fileManager, err := src.NewFileManager()
	if err != nil {
		log.Fatalln(err, "Failed to initialize FileManager")
	}

	git := src.NewGit()
	utils := src.NewUtils()

	runner := src.NewRunner(fileManager, git, utils)

	runner.Start()
}
