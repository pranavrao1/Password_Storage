package main

import (
	"encoding/json"
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
	var values []Credential
	err = json.Unmarshal([]byte(jsonOutput), &values)
	if err != nil {
		return err
	}

	cache = make(map[string]Credential)
	for _, c := range values {
		alias := c.Alias
		cache[alias] = c
	}
	return nil
}

func (cd *CredentialDatabase) locateDB() string {
	return "Database is stored at path: " + cd.filePath
}

func (cd *CredentialDatabase) searchCredentials(searchAlias string) Credential {
	return cache[searchAlias]
}
