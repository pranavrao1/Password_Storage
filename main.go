package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if !checkConfigFile() {
		fmt.Println("Please create a file called \"conf.json\" and enter the full path to store credentials.\nThe first line should be the path to json file where credentials are stored.\nSecond line should be the password to decrypt file.")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println("No inputs provided.")
		return
	}

	if len(os.Args) != 2 && len(os.Args) != 5 {
		fmt.Println("Incorrect inputs provided. Please provide a search value or four values to correctly enter new credential.")
		return
	}

	db, err := initializeDB()
	if err != nil {
		fmt.Println("Failure to initialize DB", err)
		return
	}
	if err := db.loadCredentials(); err != nil {
		fmt.Println("Failure to connect to DB", err)
		fmt.Println(db.locateDB())
		return
	}

	if len(os.Args) == 2 {
		alias := os.Args[1]
		c, ok := db.searchCredentials(alias)
		if !ok {
			fmt.Println("No credential found for alias:", alias)
			return
		}
		fmt.Println(c)
		return
	}

	if err := processCreation(db); err != nil {
		fmt.Println(err)
	}
}

func processCreation(db *CredentialDatabase) error {
	cmd := os.Args[1]
	if !strings.EqualFold(cmd, "save") {
		return fmt.Errorf("command not recognised: %s", cmd)
	}
	cred := Credential{
		Alias:    os.Args[2],
		Username: os.Args[3],
		Password: os.Args[4],
	}
	db.addCredential(cred)
	if err := db.saveCredentials(); err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}
	fmt.Println("Saved credential for alias:", cred.Alias)
	return nil
}

func checkConfigFile() bool {
	data, err := os.ReadFile("conf.json")
	if err != nil {
		return false
	}
	lines := splitConfig(data)
	if len(lines) < 2 {
		return false
	}
	return strings.HasSuffix(lines[0], ".json")
}

func initializeDB() (*CredentialDatabase, error) {
	data, err := os.ReadFile("conf.json")
	if err != nil {
		return nil, err
	}
	lines := splitConfig(data)
	if len(lines) < 2 {
		return nil, fmt.Errorf("conf.json must contain a path and a password on separate lines")
	}
	return &CredentialDatabase{filePath: lines[0], filePassword: lines[1]}, nil
}

func splitConfig(data []byte) []string {
	raw := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	lines := make([]string, 0, len(raw))
	for _, l := range raw {
		if l = strings.TrimSpace(l); l != "" {
			lines = append(lines, l)
		}
	}
	return lines
}
