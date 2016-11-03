package logger

type Logger interface {
	Name() string
	Mode() Mode
	CheckMode(Mode) bool
	Reset()

	Open() error
	Close() error

	PanicLogger
	InfoLogger
	ErrorLogger
	DebugLogger
	AccessLogger
}
