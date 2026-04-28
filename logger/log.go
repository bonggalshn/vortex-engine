package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func SetupLogging() {
	initLogging()
}

func initLogging() {
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	Warn = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}
