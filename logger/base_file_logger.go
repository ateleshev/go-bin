package logger

import (
	"fmt"
	"log"
)

type baseFileLogger struct {
	name string
	path string
	mode Mode

	loggers map[Mode]*log.Logger
}

func (this *baseFileLogger) init(name, path string, mode Mode) { // {{{
	this.name = name
	this.path = path
	this.mode = mode

	this.loggers = make(map[Mode]*log.Logger, ModesQnt)
} // }}}

func (this *baseFileLogger) register(mode Mode, logger *log.Logger) { // {{{
	this.loggers[mode] = logger
} // }}}

func (this *baseFileLogger) Name() string { // {{{
	return this.name
} // }}}

func (this *baseFileLogger) Mode() Mode { // {{{
	return this.mode
} // }}}

func (this *baseFileLogger) CheckMode(m Mode) bool { // {{{
	return this.mode&(m) != 0
} // }}}

func (this *baseFileLogger) Reset() { // {{{
	this.name = ""
	this.path = ""
	this.mode = 0

	delete(this.loggers, ModeAccess)
	delete(this.loggers, ModeError)
	delete(this.loggers, ModeDebug)
} // }}}

func (this *baseFileLogger) Panic(v ...interface{}) { // {{{
	panic(fmt.Sprint(v...))
} // }}}

func (this *baseFileLogger) Panicf(f string, v ...interface{}) { // {{{
	panic(fmt.Sprintf(f, v...))
} // }}}

func (this *baseFileLogger) Info(v ...interface{}) { // {{{
	if this.loggers[ModeInfo] != nil {
		this.loggers[ModeInfo].Print(v...)
	}
} // }}}

func (this *baseFileLogger) Infof(f string, v ...interface{}) { // {{{
	if this.loggers[ModeInfo] != nil {
		this.loggers[ModeInfo].Printf(f, v...)
	}
} // }}}

func (this *baseFileLogger) Error(v ...interface{}) { // {{{
	if this.loggers[ModeError] != nil {
		this.loggers[ModeError].Print(v...)
	}
} // }}}

func (this *baseFileLogger) Errorf(f string, v ...interface{}) { // {{{
	if this.loggers[ModeError] != nil {
		this.loggers[ModeError].Printf(f, v...)
	}
} // }}}

func (this *baseFileLogger) Debug(v ...interface{}) { // {{{
	if this.loggers[ModeDebug] != nil {
		this.loggers[ModeDebug].Print(v...)
	}
} // }}}

func (this *baseFileLogger) Debugf(f string, v ...interface{}) { // {{{
	if this.loggers[ModeDebug] != nil {
		this.loggers[ModeDebug].Printf(f, v...)
	}
} // }}}

func (this *baseFileLogger) Access(v ...interface{}) { // {{{
	if this.loggers[ModeAccess] != nil {
		this.loggers[ModeAccess].Print(v...)
	}
} // }}}

func (this *baseFileLogger) Accessf(f string, v ...interface{}) { // {{{
	if this.loggers[ModeAccess] != nil {
		this.loggers[ModeAccess].Printf(f, v...)
	}
} // }}}
