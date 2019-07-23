package main

import (
	"fmt"
)

func helpHandler() {
	fmt.Println("Available commands:")
	fmt.Printf("\tlist - list of saved sessions\n")
	fmt.Printf("\tsave session_name - save active telegram session\n")
	fmt.Printf("\trm session_name - remove saved session\n")
	fmt.Printf("\tpurge - remove ALL  saved sessions\n")
	fmt.Printf("\tback - back to previos session\n")
	fmt.Printf("\tswitch session_name - switch to another session. You should turn off your telegram client before!\n")
}
