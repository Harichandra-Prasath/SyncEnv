package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadSyncEnvFile(syncfile *SyncEnvFile) (string, error) {
	cdir, _ := os.Getwd()

	chash := hash(cdir)
	floc := fmt.Sprintf("%s/%s.json", SYNCENV_DIR, chash)

	_, err := os.Stat(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Current Directory is not added to SyncEnv!\nUse 'SyncEnv --init' first to add the current directory\n")
			return chash, fmt.Errorf("syncenv file not found")
		}
	}

	fmt.Printf("Current Directory is in SyncEnv!\n")

	data, err := os.ReadFile(floc)
	if err != nil {
		fmt.Println("Error in reading the file:", err)
		return chash, err
	}

	err = json.Unmarshal(data, syncfile)
	if err != nil {
		fmt.Println("Error in unpacking the file:", err)
		return chash, err
	}

	return floc, nil
}