package model

import "time"

type Log struct {
	Url     string
	Ip      string
	Data    time.Time
	KeyHora int
}
