package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func switchHandler() {
	if len(os.Args) == 2 {
		exitWithError("Session name is required!")
	}

	sessionPath := homedir + os.Args[2]

	if err := dirExist(sessionPath); err != nil {
		exitWithError("Session not found")
	}

	if err := dirExist(tgmpath); err != nil {
		exitWithError(err.Error())
	}

	saveLastSession()

	if err := clearDir(tgmpath); err != nil {
		exitWithError(err.Error())
	}

	err := copyRequiredFiles(tgmpath, sessionPath)
	check(err)

	fmt.Println("Success! Session switched! You can switch back to previos session with 'back' command.")
}

func saveLastSession() {
	if exist, err := neededFilesExist(tgmpath); !exist || err != nil {
		if !exist {
			return
		}

		exitWithError(err.Error())
	}

	lastSessionPath := filepath.Join(homedir, ".last_session")

	if err := dirExist(lastSessionPath); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(lastSessionPath, os.FileMode(0700))
			check(err)
		} else {
			panic(err)
		}
	}

	err := clearDir(lastSessionPath)
	check(err)

	err = copyRequiredFiles(lastSessionPath, tgmpath)
	check(err)
}
