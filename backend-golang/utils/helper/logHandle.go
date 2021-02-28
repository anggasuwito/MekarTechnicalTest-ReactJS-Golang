package helper

import (
	"log"
	"os"
	"time"
)

//LogApp is a function for create, open, and write log activity to a file
func LogApp(report string) {
	_, directoryErr := os.Stat("logactivity")
	directoryNotExist := os.IsNotExist(directoryErr)
	if directoryNotExist {
		mkdirErr := os.Mkdir("logactivity", 0755)
		if mkdirErr != nil {
			log.Print("Create logactivity directory error")
			return
		}
	}
	logFile, logFileErr := os.OpenFile("logactivity/logdata.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if logFileErr != nil {
		log.Print("Create, open, or write logdata error")
		return
	}
	timeNow := time.Now().Format(`2006-01-02 15:04:05`)
	_, writeLogDataErr := logFile.WriteString(timeNow + " | " + report + "\n")
	if writeLogDataErr != nil {
		log.Print("Logdata writestring error")
		return
	}
}
