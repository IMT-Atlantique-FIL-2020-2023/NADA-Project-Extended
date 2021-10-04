package myLog

import (
	"log"
	"os"
)

type logLevel struct {
	label string
	level int
}

var (
	//change this according to desired log level
	CURRENT_LOG_LEVEL logLevel = Get_level_INFO()
)

func Get_level_INFO() logLevel {
	return logLevel{label: "INFO", level: 0}
}

func Get_level_WARNING() logLevel {
	return logLevel{label: "WARNING", level: 1}
}

func Get_level_ERROR() logLevel {
	return logLevel{label: "ERROR", level: 2}
}

func init() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("nada-transform_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

}

func MyLog(loglevel logLevel, msg string) {
	if loglevel.level >= CURRENT_LOG_LEVEL.level {
		log.Println(" - " + loglevel.label + " - " + msg)
	}
}
