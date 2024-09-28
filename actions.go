package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

func hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// This Action adds the Current directory to SyncEnv
func InitAction() {

	cdir, _ := os.Getwd()

	// Create a json with hash of current directory

	chash := hash(cdir)
	floc := fmt.Sprintf("%s/%s.json", SYNCENV_DIR, chash)

	// Empty SyncEnv file
	syncfile := SyncEnvFile{
		Entries: make([]SyncEnvEntry, 0),
	}

	_, err := os.Stat(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Adding Current Directory to SyncEnv...\n")

			data, _ := json.Marshal(syncfile)

			// create an empty file
			err = os.WriteFile(floc, data, 0702)
			if err != nil {
				fmt.Println("Error in Creating the JSON:", err)
				return
			}

			fmt.Printf("Current Directory Added Successfully\n")
			return
		} else {
			fmt.Println("UnKnown Error in Init Action:", err)
			return
		}
	} else {
		fmt.Printf("Current Directory already added to SyncEnv!\nUse 'SyncEnv --load' to load the variables\n")
		return
	}

}

// This action unpacks the variables and set them up for loading
func unPackAction() {

	cdir, _ := os.Getwd()

	chash := hash(cdir)
	floc := fmt.Sprintf("%s/%s.json", SYNCENV_DIR, chash)

	var syncfile SyncEnvFile

	_, err := os.Stat(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Current Directory is not added to SyncEnv!\nUse 'SyncEnv --init' first to add the current directory\n")
			return
		}
	}

	fmt.Printf("Current Directory is in SyncEnv!\n")

	data, err := os.ReadFile(floc)
	if err != nil {
		fmt.Println("Error in reading the file:", err)
		return
	}

	err = json.Unmarshal(data, &syncfile)
	if err != nil {
		fmt.Println("Error in unpacking the file:", err)
		return
	}

	l := len(syncfile.Entries)

	fmt.Printf("Found %d entries\n", l)
	if l == 0 {
		return
	}
	fmt.Printf("Unpacking the Variables...\n")

	_unpack_envs(&syncfile, chash)

}

// This actions loads the latest unpacked variables
func loadAction() {

	cdir, _ := os.Getwd()

	chash := hash(cdir)
	floc := fmt.Sprintf("/tmp/%s.txt", chash)

	data, err := os.ReadFile(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("echo No file found to load. use 'SyncEnv --unpack' first")
			return
		} else {
			fmt.Println("echo 'Some Unknown Error happened'")
			return
		}
	}

	os.Remove(floc)
	fmt.Println(string(data))

}
