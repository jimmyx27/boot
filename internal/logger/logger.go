package logger

import (
	"fmt"
)

type Logger struct {
	msgs         []string
	logToConsole bool
}

func NewLogger(logToConsole bool) *Logger {
	return &Logger{
		msgs:         []string{},
		logToConsole: logToConsole,
	}
}

func (l *Logger) Flush() []string {
	msgsCopy := make([]string, len(l.msgs))
	copy(msgsCopy, l.msgs)
	l.msgs = []string{}
	return msgsCopy
}
