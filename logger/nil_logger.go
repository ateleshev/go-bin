package logger

import (
	"fmt"
)

// One nil logger
func NewNilLogger() Logger {
	return &nilLogger{}
}

type nilLogger struct {
}

func (this *nilLogger) Name() string { // {{{
	return "nil"
} // }}}

func (this *nilLogger) Mode() Mode { // {{{
	return ModeAll
} // }}}

func (this *nilLogger) CheckMode(m Mode) bool { // {{{
	return true
} // }}}

func (this *nilLogger) Reset() { // {{{
} // }}}

func (this *nilLogger) Open() (err error) { // {{{
	return
} // }}}

func (this *nilLogger) Close() error { // {{{
	return nil
} // }}}

func (this *nilLogger) Panic(v ...interface{}) { // {{{
	panic(fmt.Sprint(v...))
} // }}}

func (this *nilLogger) Panicf(f string, v ...interface{}) { // {{{
	panic(fmt.Sprintf(f, v...))
} // }}}

func (this *nilLogger) Info(v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Infof(f string, v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Error(v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Errorf(f string, v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Debug(v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Debugf(f string, v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Access(v ...interface{}) { // {{{
} // }}}

func (this *nilLogger) Accessf(f string, v ...interface{}) { // {{{
} // }}}
