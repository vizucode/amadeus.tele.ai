package database

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type internalFile struct {
	FilePath string
}

func NewInternalFile() *internalFile {
	return &internalFile{
		FilePath: "storage/",
	}
}

func (ls *internalFile) Read(fileName string) string {
	file, err := os.Open(ls.FilePath + fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	body := struct {
		Data string `json:"data"`
	}{}

	err = json.NewDecoder(file).Decode(&body)
	if err != nil {
		log.Fatal(err.Error())
	}

	return body.Data
}

func (ls *internalFile) Write(fileName string, msg string) {
	os.RemoveAll(ls.FilePath + fileName)

	data := struct {
		Data string `json:"data"`
	}{
		Data: strings.Trim(msg, ""),
	}

	body, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.WriteFile(ls.FilePath+fileName, body, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
