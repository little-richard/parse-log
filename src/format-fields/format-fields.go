package formatFields

import (
	"strconv"
	"strings"
	"time"
)

type FieldType int

const (
	Url FieldType = iota
	Ip
	Data
	Body
	Tempo
	Default
)

const (
	LabelData  string = "HORA: "
	LabelUrl   string = "URL: "
	LabelIp    string = "IP: "
	LabelBody  string = "BODY: "
	LabelTempo string = "TEMPO: "
)

func BuildTypeLineFromString(line string) FieldType {
	if strings.Contains(line, LabelData) {
		return Data
	} else if strings.Contains(line, LabelUrl) {
		return Url
	} else if strings.Contains(line, LabelIp) {
		return Ip
	} else if strings.Contains(line, LabelBody) {
		return Body
	} else if strings.Contains(line, LabelTempo) {
		return Tempo
	} else {
		return Default
	}
}

func formatLineByType(line string, typeLine FieldType) string {

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

func getCutString(typeLine FieldType) string {

	switch typeLine {
	case Ip:
		return LabelIp
	case Url:
		return LabelUrl
	case Data:
		return LabelData
	default:
		return ""
	}
}

func getDateFromLine(line string) time.Time {
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

	return dataFim
}
