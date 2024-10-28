package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// This Action adds the Current directory to SyncEnv
func InitAction() {

	cdir, _ := os.Getwd()

	// Create a SyncEnv file with hash of current directory

	chash := hash(cdir)
	floc := fmt.Sprintf("%s/%s.sy", SYNCENV_DIR, chash)

	_, err := os.Stat(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Adding Current Directory to SyncEnv...\n")

			// create an empty file
			err = os.WriteFile(floc, []byte{}, 0702)
			if err != nil {
				fmt.Println("Error in Creating the File:", err)
				return
			}

			fmt.Printf("Current Directory Added Successfully\n")
			return
		} else {
			fmt.Println("UnKnown Error in Init Action:", err)
			return
		}
	} else {
		fmt.Printf("Current Directory already added to SyncEnv!\nUse 'SyncEnv load' to load the variables\n")
		return
	}

}

// This actions loads the variables and should be eval'ed from shell
func loadAction(by_hook bool, file_path string) {

	var msg string
	var output string

	// If local file is passed
	if file_path != "" {
		data, err := _load_from_file(file_path)
		if err != nil {
			if os.IsNotExist(err) {
				msg = "echo File not found."
			} else {
				msg = "echo Some Unknown Error happened"
			}
		}
		output = data
	} else {
		cdir, _ := os.Getwd()

		chash := hash(cdir)
		floc := fmt.Sprintf("%s/%s.sy", SYNCENV_DIR, chash)

		data, err := os.ReadFile(floc)
		if err != nil {
			if os.IsNotExist(err) {
				msg = "echo No file found to load. use 'SyncEnv --init' first."
			} else {
				msg = "echo Some Unknown Error happened."
			}

		}

		output = string(data)
	}

	// Pop out the messages
	if !by_hook && msg != "" {
		fmt.Println(msg)
		return
	}

	fmt.Print(output)
}

// Action to add new variables
func addAction() {

	var syncfile SyncEnvFile

	floc, _, err := loadSyncEnvFile(&syncfile)
	if err != nil {
		return
	}

	fmt.Printf("Adding the new variables...\n")
	_add_env(floc, &syncfile)
	fmt.Printf("Addition Action Completed\n")

}

// Function to update the existing variables
func updateAction() {
	var syncfile SyncEnvFile

	floc, _, err := loadSyncEnvFile(&syncfile)
	if err != nil {
		return
	}

	fmt.Printf("Updating the requested variables...\n")
	_update_env(floc, &syncfile)
	fmt.Printf("Updation Action Completed\n")

}

// Funciton to look at the variables stored
func peekAction() {
	var syncfile SyncEnvFile

	_, _, err := loadSyncEnvFile(&syncfile)
	if err != nil {
		return
	}

	fmt.Printf("Peek results!!!\n\n")

	for _, entry := range syncfile.Entries {
		fmt.Printf("%s=%s\n", entry.Key, entry.Value)
	}
}

func _load_from_file(path string) (string, error) {

	output := ""

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		output += fmt.Sprintf("export %s\n", line)
	}

	return output, nil
}

// Action that ports the SycnEnv variables to file
func portAction() {

	file, err := os.Create(portFlag)
	if err != nil {
		fmt.Printf("Error in Creating the file: %s\n", err)
	}

	var syncfile SyncEnvFile

	_, _, err = loadSyncEnvFile(&syncfile)
	if err != nil {
		return
	}

	for _, entry := range syncfile.Entries {
		file.WriteString(fmt.Sprintf("%s=%s\n", entry.Key, entry.Value))
	}

	fmt.Println("Ported Successfully")

}

// Function to override hook the SyncEnv to current Session
func hookAction() {

	switch shellFlag {
	case "bash":
		fmt.Print(BASH_HOOK)
	case "zsh":
		fmt.Print(ZSH_HOOK)
	case "":
		fmt.Println("echo No shell provided")
	default:
		fmt.Println("echo Shell not recognised. SyncEnv supports bash and zsh")
	}

}
