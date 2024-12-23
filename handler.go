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

func handleInit() {
	InitAction(migrateFlag)
}

// Handler for Top Level Flags
func handleTop() {

	// Only one action allowed per run
	if peekFlag {
		peekAction()
	} else if helpFlag {
		fmt.Printf("%s%s\n", MAIN_TEMPLATE, LOAD_TEMPLATE)
	} else if len(addFlag) != 0 {
		addAction()
	} else if len(updateFlag) != 0 {
		updateAction()
	} else if portFlag != "" {
		portAction()
	}
}
