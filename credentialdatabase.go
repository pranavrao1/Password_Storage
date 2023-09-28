package main

import (
	"fmt"
	"os"
)

type CredentialDatabase struct {
	filePath     string
	filePassword string
}

var cache map[string]Credential

func (cd *CredentialDatabase) loadCredentials() {
	_, err := os.Open(cd.filePath)
	if err != nil {
		fmt.Println("Unable to read file provided in conf.json")
	}
}

func (cd *CredentialDatabase) dbSetup() (string, string) {
	return cd.filePath, cd.filePassword
}

func loadFile() bool {
	return true
}
