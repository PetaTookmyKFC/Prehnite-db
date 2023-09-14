package keystore

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
)

type KeyStore struct {
	ActiveSize int
	CurrSize   int

	StoreLocation string

	Storage   map[string]Data
	FirstItem string
	LastItem  string
}
type Data struct {
	Value string
	Prev  string
	Next  string
}

func Start_StringStore(size int, directory string) *KeyStore {
	rKeyStore := KeyStore{
		StoreLocation: directory,
		Storage:       make(map[string]Data),
		ActiveSize:    size,
		FirstItem:     "",
		LastItem:      "",
	}
	rKeyStore.LoadPrev()
	return &rKeyStore
}

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
func (key *KeyStore) LoadPrev() error {
	// Get Directory
	directory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	oldData := path.Join(directory, "OldMap")

	// Check if file exists
	if _, err := os.Stat(oldData); err != nil {
		return nil
	}

	file, err := os.Open(oldData)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Loaded: %s \n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return nil
}
