package main

import (
	"flag"
	"fmt"
	"os"
)

// custom multi flag
type MultiFlag []string

func (i *MultiFlag) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *MultiFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var initFlag bool
var unpackFlag bool
var loadFlag bool
var peekFlag bool

var addFlag MultiFlag
var updateFlag MultiFlag
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
			err = os.Mkdir(fmt.Sprintf("%s/packed", SYNCENV_DIR), 0702)
			if err != nil {
				panic(err)
			}
			err = os.Mkdir(fmt.Sprintf("%s/unpacked", SYNCENV_DIR), 0702)
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
	flag.BoolVar(&loadFlag, "load", false, "Flag used to load the variables")
	flag.BoolVar(&peekFlag, "peek", false, "Flag used to have a glance at stored variables")
	flag.Var(&addFlag, "add", "Flag used to add variables")
	flag.Var(&updateFlag, "update", "Flag used to update variables")

}

func main() {
	flag.Parse()

	// Only one action allowed per run
	if initFlag {
		InitAction()
	} else if unpackFlag {
		unPackAction()
	} else if loadFlag {
		loadAction()
	} else if peekFlag {
		peekAction()
	} else if len(addFlag) != 0 {
		addAction()
	} else if len(updateFlag) != 0 {
		updateAction()
	}

	fmt.Println()
}
