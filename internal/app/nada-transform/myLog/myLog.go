package myLog

import (
	"log"
	"os"
	"fmt"
	"strconv"
	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/env"
	
)

type logLevel struct {
	label string
	level int
}

var CURRENT_LOG_LEVEL logLevel

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
	logEnvLevel, err := strconv.Atoi(env.GetEnv("LOG_LEVEL"))
	if err != nil {
		MyLog(Get_level_ERROR(), err.Error())
	}
	CURRENT_LOG_LEVEL = logLevel{label: "CURRENT_LOG_LEVEL", level: logEnvLevel}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("nada-transform_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		MyLog(Get_level_ERROR(), err.Error())
	}
	log.SetOutput(file)
	
}

func MyLog(loglevel logLevel, msg string) {
	if loglevel.level >= CURRENT_LOG_LEVEL.level {
		fmt.Println(" - " + loglevel.label + " - " + msg)
		log.Println(" - " + loglevel.label + " - " + msg)
	}
	if CURRENT_LOG_LEVEL.level >= 1 {
		log.Fatal(msg)
	}
}
