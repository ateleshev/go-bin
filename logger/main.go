package logger

import (
	"os"
	"path/filepath"
)

var (
	DefaultFilenameSep string      = "-"
	DefaultFileExt     string      = ".log"
	DefaultFileFlag    int         = os.O_RDWR | os.O_CREATE | os.O_APPEND
	DefaultFilePerm    os.FileMode = 0660

	DefaultLogger = NewNilLogger()
)

// File name builders

type FilenameBuilder func(path, name string) string
type FilenameForModeBuilder func(path, name string, mode Mode) string

// Default builders

var DefaultFilenameBuilder FilenameBuilder = func(path, name string) string { // {{{
	return filepath.Join(path, name+DefaultFileExt)
} // }}}

var DefaultFilenameForModeBuilder FilenameForModeBuilder = func(path, name string, mode Mode) string { // {{{
	return filepath.Join(path, name+DefaultFilenameSep+ModeName[mode]+DefaultFileExt)
} // }}}
