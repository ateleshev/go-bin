package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	DefaultFilenameSep string      = "-"
	DefaultFileExt     string      = ".log"
	DefaultFileFlag    int         = os.O_RDWR | os.O_CREATE | os.O_APPEND
	DefaultFilePerm    os.FileMode = 0660

	DefaultLogger = NewNilLogger()
)

// File name builders

type PrefixBuilder func(Mode) string
type FilenameBuilder func(string, string) string
type FilenameForModeBuilder func(string, string, Mode) string

// Default builders

var DefaultPrefixBuilder PrefixBuilder = func(mode Mode) string { // {{{
	return fmt.Sprintf("[%d] %s: ", os.Getpid(), strings.ToUpper(ModeNames[mode]))
} // }}}

var DefaultFilenameBuilder FilenameBuilder = func(path, name string) string { // {{{
	return filepath.Join(path, name+DefaultFileExt)
} // }}}

var DefaultFilenameForModeBuilder FilenameForModeBuilder = func(path, name string, mode Mode) string { // {{{
	return filepath.Join(path, name+DefaultFilenameSep+ModeNames[mode]+DefaultFileExt)
} // }}}
