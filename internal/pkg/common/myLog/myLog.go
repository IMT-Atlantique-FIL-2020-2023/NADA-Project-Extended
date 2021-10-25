package myLog

import (
	"fmt"
	"log"
	"os"
	"strconv"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/env"
)

type Logstatus struct {
	label string
	level int
	mod   int
}

var CURRENT_LOG_LEVEL Logstatus

func Get_level_INFO() Logstatus {
	return Logstatus{label: "INFO", level: 0}
}

func Get_level_WARNING() Logstatus {
	return Logstatus{label: "WARNING", level: 1}
}

func Get_level_ERROR() Logstatus {
	return Logstatus{label: "ERROR", level: 2}
}

func Init(loglevelstring string) {
	fmt.Println(loglevelstring)
	logEnvLevel, err := strconv.Atoi(loglevelstring)
	if err != nil {
		MyLog(Get_level_ERROR(), err.Error())
	}
	logmod, err := strconv.Atoi(env.GetEnv("LOGMOD"))
	if err != nil {
		MyLog(Get_level_ERROR(), err.Error())
	}
	CURRENT_LOG_LEVEL = Logstatus{label: "CURRENT_LOG_LEVEL", level: logEnvLevel, mod: logmod}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(env.GetEnv("LOGFILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		MyLog(Get_level_ERROR(), err.Error())
	}
	log.SetOutput(file)

}

func MyLog(messageLogStatus Logstatus, msg string) {

	if CURRENT_LOG_LEVEL.mod > 0 {
		if CURRENT_LOG_LEVEL.mod > 0 && messageLogStatus.level >= CURRENT_LOG_LEVEL.level {
			if CURRENT_LOG_LEVEL.mod > 1 {
				fmt.Println(" - " + messageLogStatus.label + " - " + msg)
				if CURRENT_LOG_LEVEL.mod > 2 {
					log.Println(" - " + messageLogStatus.label + " - " + msg)
				}
			}
		}
	}
	if messageLogStatus.level >= 1 {
		log.Fatal(msg)
	}
}
