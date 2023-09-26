package main

type FileProcessor struct {
	FilePath string
	FilePassword string

}

type CredentialMetaData struct {
	metaData Credential
}

var cache map[string]CredentialMetaData

func processFile(filePath string, FilePassword string) {

}

func loadFile() bool {
	return true;
}