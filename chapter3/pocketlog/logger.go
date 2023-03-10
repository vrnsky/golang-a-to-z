package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New return you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		_, _ = fmt.Fprintf(l.output, format, args...)
	}
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Errorf formats and prints message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	_, _ = fmt.Printf(format+"\n", args)
}

func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args)
}

// logf print the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(level Level, format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, string(level)+":"+format+"\n", args...)
}
