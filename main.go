package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// func Dump() {
// 	cred1 := Credential{alias: "Alias"}
// 	fmt.Println("Running Credential", cred1)
// 	fmt.Println(ConfigExits())
// 	_ = ioutil.WriteFile("test.json", []byte(os.Args[0]), 0644)
// }

func main() {

	if !checkConfigFile() {
		fmt.Println("Please create a file called \"conf.json\" and enter the full path to store credentials.\nThe first line should be the path to json file where credentials are stored.\nSecond line should be the password to decrypt file.")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println("No inputs provided.")
		return
	}

	if len(os.Args) != 2 || len(os.Args) != 5 {
		fmt.Println("Incorrect inputs provided. Please provide a search value or four values to correctly enter new credential.")
		return
	}

	db := initializeDB()
	fmt.Println(db.dbSetup())
	db.loadCredentials()

	if len(os.Args) == 2 {
		// Search for value
		return
	}

	if len(os.Args) == 5 {
		// processCreation()
		return
	}
}

func processCreation() {
	cmd := string(os.Args[1])
	if !strings.EqualFold(cmd, "save") {
		fmt.Println("Comand not recognised: " + cmd)
	}

	cred := Credential{alias: string(os.Args[2]), username: string(os.Args[3]), password: string(os.Args[4])}
	fmt.Println("Saving Credential ", cred)
	return
}

func checkConfigFile() bool {
	_, err := os.ReadFile("conf.json")
	if err != nil {
		return false
	}

	config_file, err := os.Open("conf.json")
	if err != nil {
		return false
	}

	result, err := ioutil.ReadAll(config_file)
	if err != nil {
		return false
	}

	lines := strings.Split(string(result), "\n")
	if len(lines) != 2 {
		return false
	}

	return strings.HasSuffix(lines[0], ".json")
}

func initializeDB() *CredentialDatabase {
	config_file, _ := os.Open("conf.json")
	result, _ := ioutil.ReadAll(config_file)
	lines := strings.Split(string(result), "\n")
	return &CredentialDatabase{lines[0], lines[1]}
}
