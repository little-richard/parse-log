package main

import (
	"encoding/json"
	"io/ioutil"
	"parse-log/src/parse"
	fileUtils "parse-log/src/utils"
)

func main() {

	file, scanner := fileUtils.LoadFile("json.log")

	listaJson := parse.FileToArray(file, scanner)

	fileJson, _ := json.MarshalIndent(listaJson, "", " ")

	_ = ioutil.WriteFile("log.json", fileJson, 0644)

}
