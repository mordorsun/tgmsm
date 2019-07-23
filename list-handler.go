package main

import (
	"fmt"
	"os"
)

func listHandler() {
	dir, err := os.Open(homedir)
	check(err)

	defer dir.Close()

	files, err := dir.Readdir(0)
	check(err)

	filesCount := len(files)
	if filesCount == 0 || filesCount == 1 && files[0].Name() == ".last_session" {
		exitWithError("You have no sessions...")
	}

	i := 1
	for _, f := range files {
		if f.IsDir() && f.Name() != ".last_session" {
			fmt.Printf("%d) %s\n", i, f.Name())
			i++
		}
	}
}
