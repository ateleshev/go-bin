package logger

import (
	"log"
	"os"
	"strings"
)

// One file logger
func NewFileLogger(name, path string, mode Mode) Logger {
	l := &fileLogger{}
	l.init(name, path, mode)

	return l
}

type fileLogger struct {
	baseFileLogger

	file *os.File
}

func (this *fileLogger) init(name, path string, mode Mode) { // {{{
	this.baseFileLogger.init(name, path, mode)
} // }}}

func (this *fileLogger) Reset() { // {{{
	this.baseFileLogger.Reset()
	this.file = nil
} // }}}

func (this *fileLogger) Open() (err error) { // {{{
	this.file, err = os.OpenFile(DefaultFilenameBuilder(this.path, this.name), DefaultFileFlag, DefaultFilePerm)

	// Access
	if this.mode&(ModeAccess) != 0 {
		this.register(ModeAccess, log.New(this.file, strings.ToUpper(ModeName[ModeAccess])+": ", log.LstdFlags))
	}
	// Error
	if this.mode&(ModeError) != 0 {
		this.register(ModeError, log.New(this.file, strings.ToUpper(ModeName[ModeError])+": ", log.LstdFlags))
	}
	// Debug
	if this.mode&(ModeDebug) != 0 {
		this.register(ModeDebug, log.New(this.file, strings.ToUpper(ModeName[ModeDebug])+": ", log.LstdFlags))
	}

	return
} // }}}

func (this *fileLogger) Close() error { // {{{
	defer this.Reset()

	return this.file.Close()
} // }}}
