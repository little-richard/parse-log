package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	formatFields "parseLogPostoAki/src/format-fields"
	"time"
)

type Log struct {
	Url  string
	Ip   string
	Data time.Time
}

func main() {

	file, scanner := loadFile("json.log")

	listaJson := parseFileToArray(file, scanner)

	fileJson, _ := json.MarshalIndent(listaJson, "", " ")

	_ = ioutil.WriteFile("log.json", fileJson, 0644)

}

func loadFile(fileName string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 60*1024)
	scanner.Buffer(buf, 1024*1024)

	return file, scanner
}

func parseFileToArray(file *os.File, scanner *bufio.Scanner) []any {

	logStruct := Log{}
	lista := list.New()

	for scanner.Scan() {

		line := scanner.Text()

		fieldType := formatFields.BuildTypeLineFromString(line)

		if fieldType == formatFields.Default {
			lista.PushBack(logStruct)
			logStruct = Log{}
		}

		lineFormat := formatFields.formatLineByType(line, fieldType)

		if lineFormat != "" {

			switch fieldType {

			case Url:
				logStruct.Url = lineFormat
				break
			case Ip:
				logStruct.Ip = lineFormat
				break
			case Data:
				logStruct.Data = getDateFromLine(lineFormat)
			case Default:
				break

			}
		}

	}

	closeFile(file)

	fmt.Printf("QUANTIDADE: %v\n", lista.Len())

	array := make([]any, lista.Len())

	indice := 0

	for element := lista.Front(); element != nil; element = element.Next() {
		value := element.Value
		array[indice] = value
		indice++
	}

	return array
}

func closeFile(file *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
