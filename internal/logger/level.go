package logger

import (
	"os"
	"sync"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

var (
	logFiles   map[Level]*os.File
	levelsUsed map[Level]bool
	logChan    chan entry

	mu     sync.RWMutex
	closed bool

	wg sync.WaitGroup
)

type entry struct {
	level   Level
	message string
}
