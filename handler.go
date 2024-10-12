package main

import "fmt"

// Handler for loadFlagSet
func handleLoad() {
	loadAction(nodebugFlag, loadFromFileFlag)
}

// Handler for hookFlagSet
func handleHook() {
	hookAction()
}

// Handler for Top Level Flags
func handleTop() {

	// Only one action allowed per run
	if initFlag {
		InitAction()
	} else if unpackFlag {
		unPackAction()
	} else if peekFlag {
		peekAction()
	} else if helpFlag {
		fmt.Printf("%s%s\n", MAIN_TEMPLATE, LOAD_TEMPLATE)
	} else if len(addFlag) != 0 {
		addAction()
	} else if len(updateFlag) != 0 {
		updateAction()
	}
}
