package config

import (
	"errors"
	"log"
	"os"
	"path"

	"encoding/json"
)

/*
Only Returns error if the error is unexprected
*/
func FileExists(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

type Settings struct {
	StoreDirectory string
}

func getConfigFolder() *string {
	ucd, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	cdir := path.Join(ucd, "Prehnite")
	return &cdir
}

// Save Settings
func (config *Settings) Save() {
	// Check Directory Exists
	MakeDirectory(*getConfigFolder())

	// Create File and Object
	file, err := os.Create(path.Join(*getConfigFolder(), "base.json"))
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	file.Write(data)

	// Close File After
	defer file.Close()
}

// Set default Settings
func (config *Settings) LoadDefault() {
	// Calculate home directory
	uhd, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	config.StoreDirectory = path.Join(uhd, "Prehnite")

}

// Makes a directory if doesn't exist
// Crashes is unexprectyed error
// Returns true if item was created
// Returns false if item already exist
func MakeDirectory(directoryPath string) bool {
	if ok, err := FileExists(directoryPath); !ok {
		if err != nil {
			log.Panic(err)
		}
		// Folder Doesn't Exist ( Make it )
		err := os.Mkdir(directoryPath, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
		return true
	}
	return false
}

func LoadSystemConfig() (*Settings, error) {
	if MakeDirectory(*getConfigFolder()) {
		// Directory Didn't Exist - Create Default Config
		config := Settings{}
		config.LoadDefault()
		config.Save()
		return &config, nil
	}

	// Folder Exists
	data, err := os.ReadFile(path.Join(*getConfigFolder(), "base.json"))
	if err != nil {
		log.Fatal(err)
	}

	// Parse Json

	var config Settings
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config, nil
}
