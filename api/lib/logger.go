package lib

import (
	"fmt"
)

type (
	Logger interface {
		LogError(err error)
	}

	LoggerWithConfig struct {
		Logger
		name string
	}

	NewLoggerFunc func(name string) Logger
)

var NewLogger NewLoggerFunc = func(name string) Logger {
	return LoggerWithConfig{
		name: name,
	}
}

func (logger LoggerWithConfig) LogError(err error) {
	fmt.Printf("From %s:\n%+v", logger.name, err)
}
