package logger

type AccessLogger interface {
	Access(...interface{})
	Accessf(string, ...interface{})
}
