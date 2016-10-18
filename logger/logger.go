package logger

type Logger interface {
	Name() string
	Mode() Mode
	Reset()

	Open() error
	Close() error

	AccessLogger
	ErrorLogger
	DebugLogger
}
