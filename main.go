package main

import (
	"fmt"
	"os"
)

type Credential struct {
	alias string
	username string
	password string
}

func main() {
	cred1 := Credential{alias: "Alias"}
	fmt.Println("Running Crede", cred1)
	fmt.Println(os.Args)
}