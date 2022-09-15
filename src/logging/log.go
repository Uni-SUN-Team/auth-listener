package logging

import (
	"log"
	"os"
	"unisun/api/unisun-authen-listener/src/constants"
)

func Println(message string, error string) {
	LOG_FILE := os.Getenv(constants.LOG_PATH)
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(message, error)
}
