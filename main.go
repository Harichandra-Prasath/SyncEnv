package main

import (
	"flag"
	"fmt"
	"os"
)

var initFlag bool
var unpackFlag bool
var SYNCENV_DIR string

func init() {

	fmt.Println()

	// Check for the Directory , else create it
	SYNCENV_DIR = fmt.Sprintf("%s/.SyncEnv", os.Getenv("HOME"))

	_, err := os.Stat(SYNCENV_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(SYNCENV_DIR, 0702)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	// Define the flags
	flag.BoolVar(&initFlag, "init", false, "Flag used to add current directory to SyncEnv")
	flag.BoolVar(&unpackFlag, "unpack", false, "Flag used to unpack the variables")

}

func main() {
	flag.Parse()
	if initFlag {
		InitAction()
	} else if unpackFlag {
		unPackAction()
	}

	fmt.Println()
}
