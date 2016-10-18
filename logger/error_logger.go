package logger

type ErrorLogger interface {
	Error(...interface{})
	Errorf(string, ...interface{})
}
