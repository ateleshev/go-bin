package logger

type InfoLogger interface {
	Info(...interface{})
	Infof(string, ...interface{})
}
