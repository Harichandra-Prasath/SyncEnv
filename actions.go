package main

import (
	"crypto/sha256"
	"encoding/hex"
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

	_, err := os.Stat(floc)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Adding Current Directory to SyncEnv...\n")

			// create an empty file
			err = os.WriteFile(floc, []byte{}, 0702)
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
		fmt.Printf("Current Directory already attached to SyncEnv!\nUse 'SyncEnv --load' to load the variables\n")
		return
	}

}
