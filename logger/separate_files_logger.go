package logger

import (
	"log"
	"os"
	"strings"
)

// Use separate files to write logs by mode
func NewSeparateFilesLogger(name, path string, mode Mode) Logger {
	l := &separateFilesLogger{}
	l.init(name, path, mode)

	return l
}

type separateFilesLogger struct {
	baseFileLogger

	files map[Mode]*os.File
}

func (this *separateFilesLogger) init(name, path string, mode Mode) {
	this.baseFileLogger.init(name, path, mode)
	this.files = make(map[Mode]*os.File, ModesQnt)
}

func (this *separateFilesLogger) Reset() { // {{{
	this.baseFileLogger.Reset()

	delete(this.files, ModeAccess)
	delete(this.files, ModeError)
	delete(this.files, ModeDebug)
} // }}}

func (this *separateFilesLogger) Open() (err error) { // {{{
	// Access
	if this.mode&(ModeAccess) != 0 {
		if this.files[ModeAccess], err = os.OpenFile(DefaultFilenameForModeBuilder(this.path, this.name, ModeAccess), DefaultFileFlag, DefaultFilePerm); err != nil {
			return
		}
		this.register(ModeAccess, log.New(this.files[ModeAccess], strings.ToUpper(ModeName[ModeAccess])+": ", log.LstdFlags))
	}
	// Error
	if this.mode&(ModeError) != 0 {
		if this.files[ModeError], err = os.OpenFile(DefaultFilenameForModeBuilder(this.path, this.name, ModeError), DefaultFileFlag, DefaultFilePerm); err != nil {
			return
		}
		this.register(ModeError, log.New(this.files[ModeError], strings.ToUpper(ModeName[ModeError])+": ", log.LstdFlags))
	}
	// Debug
	if this.mode&(ModeDebug) != 0 {
		if this.files[ModeDebug], err = os.OpenFile(DefaultFilenameForModeBuilder(this.path, this.name, ModeDebug), DefaultFileFlag, DefaultFilePerm); err != nil {
			return
		}
		this.register(ModeDebug, log.New(this.files[ModeDebug], strings.ToUpper(ModeName[ModeDebug])+": ", log.LstdFlags))
	}

	return
} // }}}

func (this *separateFilesLogger) Close() (err error) { // {{{
	defer this.Reset()

	for _, file := range this.files {
		if err = file.Close(); err != nil {
			return
		}
	}

	return
} // }}}
