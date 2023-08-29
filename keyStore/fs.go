package keystore

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
)

// Create item that contains the keyValue that was loaded in RAM
func (key *KeyStore) Save() error {
	// Get Directory
	directory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	oldData := path.Join(directory, "OldMap")
	// Create file
	file, err := os.Create(oldData)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	// Write Each line to buffer
	for value := range key.Storage {
		w.WriteString(fmt.Sprintf("%s\n", value))
	}

	// Write Buffer to file
	return w.Flush()
}

// Load file
// func (key *KeyStore) LoadPrev() error {
// 	// Get Directory
// 	directory, err := os.Getwd()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	oldData := path.Join(directory, "OldMap")

// 	return nil
// }
