package formatFields

import (
	enum "parse-log/model/enums"
	label "parse-log/strings"
	"strconv"
	"strings"
	"time"
)

func BuildTypeLineFromString(line string) enum.FieldType {
	if strings.Contains(line, label.Data) {
		return enum.Data
	} else if strings.Contains(line, label.Url) {
		return enum.Url
	} else if strings.Contains(line, label.Ip) {
		return enum.Ip
	} else if strings.Contains(line, label.Body) {
		return enum.Body
	} else if strings.Contains(line, label.Tempo) {
		return enum.Tempo
	} else {
		return enum.Default
	}
}

func FormatLineByType(line string, typeLine enum.FieldType) string {

	cutString := getCutString(typeLine)

	if cutString == "" {
		return ""
	}

	_, after, found := strings.Cut(line, cutString)

	if found {
		return after
	} else {
		return ""
	}
}

func getCutString(typeLine enum.FieldType) string {

	switch typeLine {
	case enum.Ip:
		return label.Ip
	case enum.Url:
		return label.Url
	case enum.Data:
		return label.Data
	default:
		return ""
	}
}

func GetDateFromLine(line string) (time.Time, int) {
	dataHoraSplit := strings.Split(line, " ")
	data := dataHoraSplit[0]
	horario := dataHoraSplit[1]

	dataSplit := strings.Split(data, "/")
	horarioSplit := strings.Split(horario, ":")

	dia, _ := strconv.Atoi(dataSplit[0])
	mes, _ := strconv.Atoi(dataSplit[1])
	ano, _ := strconv.Atoi(dataSplit[2])

	hora, _ := strconv.Atoi(horarioSplit[0])
	minutos, _ := strconv.Atoi(horarioSplit[1])
	segundos, _ := strconv.Atoi(horarioSplit[2])

	loc, _ := time.LoadLocation("America/Sao_Paulo")

	dataFim := time.Date(ano, time.Month(mes), dia, hora, minutos, segundos, 0, loc)

	return dataFim, hora
}
