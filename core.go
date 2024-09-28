package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type SyncEnvFile struct {
	Entries []SyncEnvEntry `json:"entries"`
}

type SyncEnvEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Iterate thorugh the entries and load it
func _unpack_envs(file *SyncEnvFile, hash string) {

	var unpacked string

	for _, entry := range file.Entries {
		unpacked += fmt.Sprintf("export %s='%s'\n", entry.Key, entry.Value)
	}

	// write it to the disk
	floc := fmt.Sprintf("/tmp/%s.txt", hash)
	os.WriteFile(floc, []byte(unpacked), 0702)

}

// Append the new entry to the SyncEnv file and write it
func _add_env(floc string, file *SyncEnvFile) {

	for _, item := range addFlag {
		_entry := strings.Split(item, "=")

		if len(_entry) != 2 {
			fmt.Println("Improper key-value Pair...")
			continue
		}

		// create a entry
		entry := SyncEnvEntry{
			Key:   _entry[0],
			Value: _entry[1],
		}

		file.Entries = append(file.Entries, entry)
	}

	// marshall it
	data, _ := json.Marshal(file)

	os.WriteFile(floc, data, 0702)

}
