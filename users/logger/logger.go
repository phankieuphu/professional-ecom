package logger

import (
	"log"
	"os"
)

type ConsoleLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO:  ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN:  ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		fatal: log.New(os.Stderr, "FATAL: ", log.LstdFlags),
	}
}

func (l *ConsoleLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *ConsoleLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *ConsoleLogger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *ConsoleLogger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *ConsoleLogger) Fatal(v ...interface{}) {
	l.fatal.Fatalln(v...)
}
