package main

type Credential struct {
	Alias    string
	Username string
	Password string
}

type CredentialCache struct {
	credentials []Credential
}
