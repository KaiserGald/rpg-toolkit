// Package logger
// 3 January, 2017
// Code is licensed under the MIT License

package logger

import (
	"io"
	"log"
)

// Logger struct
type Logger struct {
	Debug *log.Logger

	Info *log.Logger

	Error *log.Logger
}

// Init initializes the logger
func (l *Logger) Init(debugHandle io.Writer, infoHandle io.Writer, errorHandle io.Writer) {
	l.Debug = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
