package main

import (
	"flag"
	"fmt"
	"os"
)

// Custom MultiFlags for add,update,etc..
type MultiFlag []string

func (i *MultiFlag) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *MultiFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	// core dir for the SyncEnv
	SYNCENV_DIR string

	peekFlag         bool
	helpFlag         bool
	loadFromFileFlag string
	portFlag         string
	addFlag          MultiFlag
	updateFlag       MultiFlag

	// FlagSets
	loadFlagSet *flag.FlagSet
	nodebugFlag bool

	hookFlagSet *flag.FlagSet
	shellFlag   string

	initFlagSet *flag.FlagSet
	migrateFlag string
)

func init() {

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

	// Define the top level flags
	flag.BoolVar(&peekFlag, "peek", false, "Flag used to have a glance at stored variables")
	flag.BoolVar(&helpFlag, "help", false, "Flag used to show the help menu")
	flag.StringVar(&migrateFlag, "with-file", "", "Flag used to add the local file on init")
	flag.Var(&addFlag, "add", "Flag used to add variables")
	flag.Var(&updateFlag, "update", "Flag used to update variables")
	flag.StringVar(&portFlag, "port", "", "Flag used to port SyncEnv variables to file")

	// Define the FlagSets
	loadFlagSet = flag.NewFlagSet("load", flag.ExitOnError)
	hookFlagSet = flag.NewFlagSet("hook", flag.ExitOnError)
	initFlagSet = flag.NewFlagSet("init", flag.ExitOnError)

	loadFlagSet.BoolVar(&nodebugFlag, "no-debug", false, "Flag used to output messages on load action")
	loadFlagSet.StringVar(&loadFromFileFlag, "from-file", "", "Flag used to load variables from local .env file")

	hookFlagSet.StringVar(&shellFlag, "shell", "", "User's shell for the hook")

	initFlagSet.StringVar(&migrateFlag, "with-file", "", "Flag to add the local variables on init")

}

func main() {

	if len(os.Args) < 2 {
		panic("No args or flags passed")
	}

	switch os.Args[1] {
	case "init":
		initFlagSet.Parse(os.Args[2:])
		handleInit()
	case "load":
		loadFlagSet.Parse(os.Args[2:])
		handleLoad()
	case "hook":
		hookFlagSet.Parse(os.Args[2:])
		handleHook()
	default:
		flag.Parse()
		handleTop()
	}

}
