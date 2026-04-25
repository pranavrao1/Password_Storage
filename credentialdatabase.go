package main

import (
	"encoding/json"
	"os"
)

type CredentialDatabase struct {
	filePath     string
	filePassword string
	cache        map[string]Credential
}

func (cd *CredentialDatabase) loadCredentials() error {
	cd.cache = make(map[string]Credential)

	data, err := os.ReadFile(cd.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}

	jsonBytes := data
	if isEncryptedVault(data) {
		jsonBytes, err = decryptVault(data, cd.filePassword)
		if err != nil {
			return err
		}
	}

	var values []Credential
	if err := json.Unmarshal(jsonBytes, &values); err != nil {
		return err
	}
	for _, c := range values {
		cd.cache[c.Alias] = c
	}
	return nil
}

func (cd *CredentialDatabase) saveCredentials() error {
	values := make([]Credential, 0, len(cd.cache))
	for _, c := range cd.cache {
		values = append(values, c)
	}
	jsonBytes, err := json.MarshalIndent(values, "", "  ")
	if err != nil {
		return err
	}
	encrypted, err := encryptVault(jsonBytes, cd.filePassword)
	if err != nil {
		return err
	}
	return os.WriteFile(cd.filePath, encrypted, 0600)
}

func (cd *CredentialDatabase) addCredential(c Credential) {
	cd.cache[c.Alias] = c
}

func (cd *CredentialDatabase) locateDB() string {
	return "Database is stored at path: " + cd.filePath
}

func (cd *CredentialDatabase) searchCredentials(alias string) (Credential, bool) {
	c, ok := cd.cache[alias]
	return c, ok
}
