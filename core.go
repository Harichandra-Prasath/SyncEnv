package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type SyncEnvFile struct {
	Entries []*SyncEnvEntry
}

type SyncEnvEntry struct {
	Key   string
	Value string
}

// Append the new entry to the SyncEnv file and write it
func _add_env(floc string, file *SyncEnvFile) {

	for _, item := range addFlag {
		_entry := strings.SplitN(item, "=", 2)

		if len(_entry) != 2 {
			fmt.Println("Improper key-value Pair...")
			continue
		}

		// create a entry
		entry := SyncEnvEntry{
			Key:   _entry[0],
			Value: _entry[1],
		}

		file.Entries = append(file.Entries, &entry)
	}

	_write_to_sy(floc, file)

}

// Update the varibles with new value to the SyncEnv file and write it
func _update_env(floc string, file *SyncEnvFile) {

	for _, item := range updateFlag {

		// Split with first "=" character
		_entry := strings.SplitN(item, "=", 2)

		if len(_entry) != 2 {
			fmt.Println("Improper key-value Pair...")
			continue
		}

		key := _entry[0]
		new_value := _entry[1]

		_look_up_and_set(file, key, new_value)

	}

	_write_to_sy(floc, file)

}

func _look_up_and_set(file *SyncEnvFile, key string, new_value string) {

	for _, entry := range file.Entries {
		if entry.Key == key {
			entry.Value = new_value
			return
		}
	}

	fmt.Printf("Specified variable '%s' was not found\n", key)
}

func _write_to_sy(floc string, sycnfile *SyncEnvFile) {

	buffer := &bytes.Buffer{}

	for _, entry := range sycnfile.Entries {
		_entry := fmt.Sprintf("export %s=%s\n", entry.Key, entry.Value)
		buffer.WriteString(_entry)
	}

	os.WriteFile(floc, buffer.Bytes(), 0702)

}
