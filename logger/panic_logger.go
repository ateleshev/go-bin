package logger

type PanicLogger interface {
	Panic(...interface{})
	Panicf(string, ...interface{})
}
