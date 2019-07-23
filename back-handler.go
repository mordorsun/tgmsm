package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func backHandler() {
	path := homedir + "/.last_session"

	if err := dirExist(path); err != nil {
		if os.IsNotExist(err) {
			exitWithError("Last session don't exist!")
		}

		check(err)
	}

	if filesExist, err := neededFilesExist(path); !filesExist || err != nil {
		exitWithError("Last session files not found!")
	}

	exist, err := neededFilesExist(tgmpath)
	check(err)

	temppath := filepath.Join(path, "temp")
	if exist {
		err = os.Mkdir(temppath, os.FileMode(0700))
		check(err)

		err = copyRequiredFiles(temppath, tgmpath)
		check(err)
	}

	if err = clearDir(tgmpath); err != nil {
		panic(err)
	}

	if err = copyRequiredFiles(tgmpath, path); err != nil {
		panic(err)
	}

	if exist {
		for _, f := range files {
			err = os.RemoveAll(filepath.Join(path, f))
			check(err)
		}

		err = copyRequiredFiles(path, temppath)
		check(err)

		err = os.RemoveAll(temppath)
		check(err)
	}

	fmt.Println("Success! Last session returned!")
}
