package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CredentialDatabase struct {
	filePath     string
	filePassword string
}

var cache map[string]Credential

func (cd *CredentialDatabase) loadCredentials() error {
	dbFile, err := os.Open(cd.filePath)
	if err != nil {
		return err
	}
	jsonOutput, err := ioutil.ReadAll(dbFile)
	if err != nil {
		return err
	}
	var cache Credential
	err = json.Unmarshal([]byte(jsonOutput), &cache)
	if err != nil {
		return err
	}

	fmt.Println(cache)
	return nil
}

func (cd *CredentialDatabase) locateDB() string {
	return "Database is stored at path: " + cd.filePath
}

func loadFile() bool {
	return true
}
