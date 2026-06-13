package util

import (
	"log"
	"time"
)

const layout = "2006-01-02T15:04:05"

func StringToTime(value string) time.Time {
	convertedTime, err := time.Parse(layout, value)
	if err != nil {
		log.Printf("erro ao fazer parse da data: %v", err)
	}
	return convertedTime
}
