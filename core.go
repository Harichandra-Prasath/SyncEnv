package main

import (
	"fmt"
	"os"
)

type SyncEnvFile struct {
	Entries []SyncEnvEntry `json:"entries"`
}

type SyncEnvEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Iterate thorugh the entries and load it
func _load_envs(file *SyncEnvFile) {

	for _, entry := range file.Entries {
		err := os.Setenv(entry.Key, entry.Value)
		if err != nil {
			fmt.Printf("Error in loading the key-%s: %s\n", entry.Key, err)
		}
	}

}
