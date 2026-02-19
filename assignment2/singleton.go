package assignment2

import (
	"log"
	"sync"
)

type Logger struct{}

var instance *Logger
var once sync.Once

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Log(message string) {
	log.Println(message)
}
