package pocketlog

import "io"

// Option defines a functional option to our logger.
type Option func(l *Logger)

func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.writer = output
	}
}
