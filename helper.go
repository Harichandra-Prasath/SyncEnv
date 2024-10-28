package main

import (
	"fmt"
	"os"
	"strings"
)

func loadSyncEnvFile(syncfile *SyncEnvFile) (string, string, error) {
	cdir, _ := os.Getwd()
	chash := hash(cdir)
	floc := fmt.Sprintf("%s/%s.sy", SYNCENV_DIR, chash)

	_, err := os.Stat(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Current Directory is not added to SyncEnv!\nUse 'SyncEnv --init' first to add the current directory\n")
			return floc, chash, fmt.Errorf("syncenv file not found")
		}
	}

	fmt.Printf("Current Directory is in SyncEnv!\n")

	data, err := os.ReadFile(floc)
	if err != nil {
		fmt.Println("Error in reading the file:", err)
		return "", "", err
	}

	for _, entry := range strings.Split(string(data), "\n") {

		if entry != "" {

			// Strip the "export"
			entry = strings.TrimPrefix(entry, "export ")

			items := strings.SplitN(entry, "=", 2)

			syncfile.Entries = append(syncfile.Entries, &SyncEnvEntry{
				Key:   items[0],
				Value: items[1],
			})
		}

	}

	return floc, chash, nil
}
