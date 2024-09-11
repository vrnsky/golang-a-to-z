package pocketlog

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors.
	LevelError
)

type Level byte

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelError:
		return "[ERROR]"
	default:
		return ""
	}
}
