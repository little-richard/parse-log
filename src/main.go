package main

import (
	. "github.com/ahmetb/go-linq/v3"
	"log"
	"parse-log/model"
	"parse-log/parse"
	fileUtils "parse-log/utils"
)

type AgrupamentoPorIp struct {
	Ip                 string
	AgrupamentoPorHora []interface{}
}

func main() {

	log.Println("Iniciando Parse do log ")

	file, scanner := fileUtils.LoadFile("json.log")

	array := parse.FileToArray(file, scanner)

	groupByIp := From(array).GroupBy(
		func(log interface{}) interface{} {
			return log.(model.Log).Ip
		}, func(log interface{}) interface{} {
			return log
		}).Results()

	groupByHour := From(groupByIp).GroupBy(
		func(group interface{}) interface{} {
			return group.(Group).Key
		}, func(group interface{}) interface{} {
			return From(group.(Group).Group).GroupBy(
				func(log interface{}) interface{} {
					return log.(model.Log).KeyHora
				}, func(groupHour interface{}) interface{} {
					return groupHour
				}).Results()
		}).Results()

	var groupFormatado []AgrupamentoPorIp

	From(groupByHour).ForEach(func(element interface{}) {
		item := element.(Group)
		grupoHora := item.Group[0].([]interface{})

		groupFormatado = append(groupFormatado, AgrupamentoPorIp{Ip: item.Key.(string), AgrupamentoPorHora: grupoHora})
	})

	fileUtils.ParseJSON(groupByIp, "log_group_by_ip.json")
	fileUtils.ParseJSON(groupByHour, "log_group_by_ip_by_hora.json")
	fileUtils.ParseJSON(groupFormatado, "log_group_by_ip_by_hora_formatado.json")

	log.Println("Fim Parse do log para arquivo JSON")
}
