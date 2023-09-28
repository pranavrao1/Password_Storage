package main

import (
	"os"
)

type CredentialDatabase struct {
	filePath     string
	filePassword string
}

var cache map[string]Credential

func (cd *CredentialDatabase) loadCredentials() error {
	_, err := os.Open(cd.filePath)
	return err
}

func (cd *CredentialDatabase) locateDB() string {
	return "Database is stored at path: " + cd.filePath
}

func loadFile() bool {
	return true
}
