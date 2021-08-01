package logger

import (
	"fmt"
	"strings"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
}

var gLogger Logger

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

var level2Str = map[Level]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
}

func (l Level) String() string {

	return level2Str[l]
}

type logger struct {
	level Level
}

func (l logger) log(actualLevel Level, args ...interface{}) {
	if actualLevel < l.level {
		return
	}
	fmt.Println(fmt.Sprintf(strings.TrimSuffix(fmt.Sprintln(args...), "\n")))
}

func (l logger) Debug(args ...interface{}) {
	l.log(DEBUG, args...)
}

func (l logger) Info(args ...interface{}) {
	l.log(INFO, args...)
}

func (l logger) Warning(args ...interface{}) {
	l.log(WARNING, args...)
}

func (l logger) Error(args ...interface{}) {
	l.log(ERROR, args...)
}

func Init(level Level) {

	gLogger = &logger{
		level: level,
	}
}

func Instance() Logger {
	return gLogger
}
