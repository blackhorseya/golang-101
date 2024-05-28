package singleton

import (
	"sync"
)

// Logger is a struct that will be used to log messages.
type Logger struct {
}

var logger *Logger
var once sync.Once

// GetLogger returns a pointer to a Logger instance.
func GetLogger() *Logger {
	once.Do(func() {
		logger = &Logger{}
	})

	return logger
}
