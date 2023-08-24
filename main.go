package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func ConfigExits() bool {
	if _, err := os.Stat("conf.json"); err != nil {
		return false
	} else {
		return true
	}
}

func Dump() {
	cred1 := Credential{alias: "Alias"}
	fmt.Println("Running Credential", cred1)
	fmt.Println(ConfigExits())
	_ = ioutil.WriteFile("test.json", []byte(os.Args[0]), 0644)
}

type Credential struct {
	alias string
	username string
	password string
}

func main() {

	if !ConfigExits() {
		fmt.Println("Please create a file called \"conf.json\" and enter the full path to store credentials")
		return
	} 

	if len(os.Args) == 1 {
		fmt.Println("No inputs provided.")
		return
	}
}