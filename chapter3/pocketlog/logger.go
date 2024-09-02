package pocketlog

// Logger is used to log information.
type Logger struct {
	threshold Level
}

// Debugf formats and print message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {

}

// Infof formats and print message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {

}

// Errorf formats and print message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {

}

func New(threshold Level) *Logger {
	return &Logger {
		threshold: threshold,
	}
}