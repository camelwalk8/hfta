package main

import (
	"fmt"
	"log"
	"os"
)

var (
	DL *log.Logger
	IL *log.Logger
	EL *log.Logger
	FL *log.Logger
)

var LogDirectory string

func initLogging(logFileDir string) (bool, error) {
	if b, err := Exists(logFileDir); !b {
		return false, err
	}

	logFilePath := fmt.Sprintf("%v/hfta.log", logFileDir)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false, err
	}

	IL = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	DL = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	EL = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FL = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	return true, nil

}

func SetupLogging(logFileDir string) (bool, error) {

	if b, err := initLogging(logFileDir); !b {
		return false, err
	}
	Debugl("Application Logging is now setup!")
	LogDirectory = logFileDir
	return true, nil
}

// type LogDir struct {
// 	LogDirectory string
// }

// func GetLogger(logFileDir string) *LogDir {

// 	if b, _ := Exists(logFileDir); !b {
// 		err := os.Mkdir(logFileDir, 0666)
// 		if err != nil {
// 			log.Fatalln("Unable to set logging", err)
// 		}
// 	}

// 	return &LogDir{
// 		LogDirectory: logFileDir,
// 	}
// }

// func SetupLogFile(logDirPath string) *os.File {
// 	f, e := os.OpenFile(logDirPath+"/"+"hfta.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// 	return f
// }

// func (l *LogDir) Info() *log.Logger {
// 	f := SetupLogFile(l.LogDirectory)
// 	return log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func (l *LogDir) Warning() *log.Logger {
// 	f := SetupLogFile(l.LogDirectory)
// 	return log.New(f, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func (l *LogDir) Error() *log.Logger {
// 	f := SetupLogFile(l.LogDirectory)
// 	return log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func (l *LogDir) Fatal() *log.Logger {
// 	f := SetupLogFile(l.LogDirectory)
// 	return log.New(f, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
// }
