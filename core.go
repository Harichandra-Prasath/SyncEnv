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
func _unpack_envs(file *SyncEnvFile, hash string) {

	var unpacked string

	for _, entry := range file.Entries {
		unpacked += fmt.Sprintf("export %s='%s'\n", entry.Key, entry.Value)
	}

	// write it to the disk
	floc := fmt.Sprintf("/tmp/%s.txt", hash)
	os.WriteFile(floc, []byte(unpacked), 0702)

}
