package parse

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	formatFields "parse-log/src/format-fields"
	"parse-log/src/model"
	enum "parse-log/src/model/enums"
	fileUtils "parse-log/src/utils"
)

func FileToArray(file *os.File, scanner *bufio.Scanner) []any {

	logStruct := model.Log{}
	lista := list.New()

	for scanner.Scan() {

		line := scanner.Text()

		fieldType := formatFields.BuildTypeLineFromString(line)

		if fieldType == enum.Default {
			lista.PushBack(logStruct)
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
				logStruct.Data = formatFields.GetDateFromLine(lineFormat)
			case enum.Default:
				break

			}
		}

	}

	fileUtils.CloseFile(file)

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
