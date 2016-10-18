package logger

// (Error | Debug) level
const (
	LevelSilent = 1 << iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelVerbose
	LevelTiming
	LevelDebug
	LevelTrace
)
