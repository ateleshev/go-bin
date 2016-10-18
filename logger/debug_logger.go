package logger

type DebugLogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
