package main

import (
	"os"
)

var commands = map[string]func(){
	"list":   listHandler,
	"save":   saveHandler,
	"rm":     removeHandler,
	"switch": switchHandler,
	"back":   backHandler,
	"help":   helpHandler,
	"purge":  purgeHandler,
}

var tgmpath = ""
var homedir = ""

func init() {
	usrhome, err := os.UserHomeDir()
	check(err)

	homedir = usrhome + "/.local/share/tgsm/"

	if err := dirExist(homedir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(homedir, filemode)
			check(err)
		} else {
			exitWithError(err.Error())
		}
	}

	tgmpath = usrhome + "/.local/share/TelegramDesktop/tdata/"
}

func main() {
	unkCmdError := "unknown command\nwrite 'help' to get help page"

	if len(os.Args) == 1 {
		exitWithError(unkCmdError)
	}

	cmd := commands[os.Args[1]]

	if cmd == nil {
		exitWithError(unkCmdError)
	}

	cmd()
}
