package main

import (
	"fmt"
	"os"
)

func saveHandler() {
	if len(os.Args) == 2 {
		exitWithError("ERROR! Session name expected!")
	}

	if os.Args[2] == ".last_session" {
		exitWithError(".last_session is busy by this tool!")
	}

	sessionDir := homedir + os.Args[2]

	if _, err := os.Stat(sessionDir); !os.IsNotExist(err) {
		fmt.Println("This session name already exist!")
		return
	}

	f, err := os.Open(tgmpath)
	check(err)

	defer f.Close()

	if exist, err := neededFilesExist(tgmpath); !exist || err != nil {
		if !exist {
			exitWithError("Required session files not found!")
		}

		exitWithError(err.Error())
	}

	err = os.Mkdir(sessionDir, os.FileMode(0700))
	check(err)

	copyRequiredFiles(sessionDir, tgmpath)

	fmt.Println("Success! Session was saved!")
}
