package logger

type Mode uint8

// Modes of loggers: (Access | Error | Debug)
const (
	// modes quantity
	ModesQnt = 4

	ModeNone Mode = 0

	ModeInfo Mode = 1 << iota
	ModeError
	ModeDebug
	ModeAccess

	ModeAll       Mode = ModeInfo | ModeError | ModeDebug | ModeAccess
	ModeInfoError Mode = ModeInfo | ModeError
)

var ModeNames map[Mode]string = map[Mode]string{
	ModeInfo:   "info",
	ModeError:  "error",
	ModeDebug:  "debug",
	ModeAccess: "access",
}
