package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestConsoleLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	l := &ConsoleLogger{
		debug: log.New(&buf, "DEBUG: ", 0),
	}

	l.Debug("Test debug message")
	output := buf.String()

	if !strings.Contains(output, "DEBUG: Test debug message") {
		t.Errorf("expected debug log to contain message, got: %s", output)
	}
}

func TestConsoleLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	l := &ConsoleLogger{
		info: log.New(&buf, "INFO: ", 0),
	}

	l.Info("Test info message")
	output := buf.String()

	if !strings.Contains(output, "INFO: Test info message") {
		t.Errorf("expected info log to contain message, got: %s", output)
	}
}

func TestConsoleLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	l := &ConsoleLogger{
		warn: log.New(&buf, "WARN: ", 0),
	}

	l.Warn("Test warn message")
	output := buf.String()

	if !strings.Contains(output, "WARN: Test warn message") {
		t.Errorf("expected warn log to contain message, got: %s", output)
	}
}

func TestConsoleLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	l := &ConsoleLogger{
		error: log.New(&buf, "ERROR: ", 0),
	}

	l.Error("Test error message")
	output := buf.String()

	if !strings.Contains(output, "ERROR: Test error message") {
		t.Errorf("expected error log to contain message, got: %s", output)
	}
}
