package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var logFile *os.File

func Init() error {
	err := os.MkdirAll("log", 0755)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := filepath.Join("log", fmt.Sprintf("code2doc_%s.log", timestamp))

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	logFile = file
	Log("INFO", "Logger initialized")
	return nil
}

func Log(level string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, message)

	if logFile != nil {
		logFile.WriteString(logLine)
	}
	logFile.Sync()
}

func Close() {
	if logFile != nil {
		Log("INFO", "Logger closed")
		logFile.Close()
	}
}