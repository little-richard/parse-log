package utils

import (
	"bufio"
	"log"
	"os"
)

func LoadFile(fileName string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 60*1024)
	scanner.Buffer(buf, 1024*1024)

	return file, scanner
}

func CloseFile(file *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
