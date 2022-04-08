package parse

import (
	"bufio"
	"log"
	"os"
	formatFields "parse-log/format-fields"
	"parse-log/model"
	enum "parse-log/model/enums"
	fileUtils "parse-log/utils"
)

func FileToArray(file *os.File, scanner *bufio.Scanner) []model.Log {

	log.Println("Iniciando parse do arquivo para array ")

	logStruct := model.Log{}
	var array []model.Log

	for scanner.Scan() {

		line := scanner.Text()

		fieldType := formatFields.BuildTypeLineFromString(line)

		if fieldType == enum.Default {
			array = append(array, logStruct)
			logStruct = model.Log{}
		}

		lineFormat := formatFields.FormatLineByType(line, fieldType)

		if lineFormat != "" {

			switch fieldType {

			case enum.Url:
				logStruct.Url = lineFormat
				break
			case enum.Ip:
				logStruct.Ip = lineFormat
				break
			case enum.Data:
				logStruct.Data, logStruct.KeyHora = formatFields.GetDateFromLine(lineFormat)
			case enum.Default:
				break

			}
		}

	}

	log.Println("Fechando arquivo ")

	fileUtils.CloseFile(file)

	log.Printf("Fim Parse arquivo para Array, a quantidade de registros é: %v\n", len(array))

	return array
}
