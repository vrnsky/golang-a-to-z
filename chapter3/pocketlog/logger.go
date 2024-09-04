package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	writer io.Writer
}

// Debugf formats and print message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l == nil {
		l.writer = os.Stdout
	}
	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)
}

// Infof formats and print message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l == nil {
		l.writer = os.Stdout
	}
	if l.threshold > LevelInfo {
		return
	}
	l.logf(format, args...)
}

// Errorf formats and print message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.writer == nil {
		l.writer = os.Stdout
	}
	if l.threshold > LevelError {
		return
	}
	l.logf(format, args)
}

// logf prints message to output
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.writer, format + "\n", args...)
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, output io.Writer, options ...Option) *Logger {
	lgr := &Logger { threshold: threshold, writer: output}

	for _, configFunc := range options {
		configFunc(lgr)
	}

	return lgr
}