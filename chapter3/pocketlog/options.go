package pocketlog

import "io"

type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
