package main

import (
	"fmt"
	"os"
	"strings"
)

func purgeHandler() {
	dir, err := os.Open(homedir)
	check(err)

	defer dir.Close()

	files, err := dir.Readdir(0)
	check(err)

	for _, f := range files {
		n := f.Name()

		if strings.Contains(n, ".last_session") {
			continue
		}

		err = os.RemoveAll(homedir + n)
		check(err)

		fmt.Printf("Session '%s' was removed\n", n)
	}

	fmt.Println("Success! All sessions removed!")
}
