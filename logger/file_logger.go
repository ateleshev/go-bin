package logger

import (
	"log"
	"os"
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

	// Info
	if this.mode&(ModeInfo) != 0 {
		this.register(ModeInfo, log.New(this.file, DefaultPrefixBuilder(ModeInfo), log.LstdFlags))
	}
	// Error
	if this.mode&(ModeError) != 0 {
		this.register(ModeError, log.New(this.file, DefaultPrefixBuilder(ModeError), log.LstdFlags))
	}
	// Debug
	if this.mode&(ModeDebug) != 0 {
		this.register(ModeDebug, log.New(this.file, DefaultPrefixBuilder(ModeDebug), log.LstdFlags))
	}
	// Access
	if this.mode&(ModeAccess) != 0 {
		this.register(ModeAccess, log.New(this.file, DefaultPrefixBuilder(ModeAccess), log.LstdFlags))
	}

	return
} // }}}

func (this *fileLogger) Close() error { // {{{
	return this.file.Close()
} // }}}
