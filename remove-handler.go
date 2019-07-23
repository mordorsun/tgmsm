package main

import (
	"fmt"
	"os"
)

func removeHandler() {
	if len(os.Args) == 2 {
		exitWithError("Session name required!")
	}

	if os.Args[2] == ".last_session" {
		exitWithError("You can not remove system folder!")
	}

	path := homedir + os.Args[2]

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		exitWithError("Session not found")
	}

	check(err)

	f.Close()

	err = os.RemoveAll(path)
	check(err)

	fmt.Printf("Success! Session '%s' was removed!\n", os.Args[2])
}
