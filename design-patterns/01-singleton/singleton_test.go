package singleton

import (
	"testing"
)

func TestGetLoggerReturnsSingletonInstance(t *testing.T) {
	logger1 := GetLogger()
	logger2 := GetLogger()

	if logger1 != logger2 {
		t.Errorf("GetLogger() should always return the same instance")
	}
}

func TestGetLoggerConcurrentAccess(t *testing.T) {
	c := make(chan *Logger, 2)

	go func() { c <- GetLogger() }()
	go func() { c <- GetLogger() }()

	logger1 := <-c
	logger2 := <-c

	if logger1 != logger2 {
		t.Errorf("GetLogger() should always return the same instance, even in concurrent scenarios")
	}
}
