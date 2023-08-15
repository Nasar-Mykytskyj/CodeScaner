package logger

import (
	"log"
	"os"
	"sync"
)

const (
	LogsPath = "logs"
)

var generalLogger *log.Logger
var once sync.Once

func GetGeneralLogger() *log.Logger {
	once.Do(func() {
		generaLogFile := SetLogFile(LogsPath)
		generalLogger = log.New(generaLogFile, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	})

	return generalLogger
}

func SetLogFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error opening file:", err)
		os.Exit(1)
	}

	return file
}
