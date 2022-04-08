package utils

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var dirFiles = "assets"

func LoadFile(fileName string) (*os.File, *bufio.Scanner) {

	filePath := dirFiles + "/" + fileName

	log.Printf("Carregando arquivo %s ......... \n", fileName)

	file, err := os.Open(filePath)

	log.Printf("Arquivo aberto %s \n", fileName)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Criando buffer scanner ......... ")

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 60*1024)
	scanner.Buffer(buf, 1024*1024)

	log.Println("Arquivo preparado com sucesso ")

	return file, scanner
}

func CloseFile(file *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}

func ParseJSON(target any, fileOutputName string) {

	fileJson, _ := json.MarshalIndent(target, "", " ")

	outputPath := dirFiles + "/" + fileOutputName

	_ = ioutil.WriteFile(outputPath, fileJson, 0644)
}
