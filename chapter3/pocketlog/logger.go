package pocketlog

import (
	"fmt"
	"io"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New return you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
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
