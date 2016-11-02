package logger

type Mode uint8

// Modes of loggers: (Access | Error | Debug)
const (
	// modes quantity
	ModesQnt = 4

	ModeInfo Mode = 1 << iota
	ModeError
	ModeDebug
	ModeAccess

	ModeAll       Mode = ModeInfo | ModeError | ModeDebug | ModeAccess
	ModeInfoError Mode = ModeInfo | ModeError
)

var ModeName map[Mode]string = map[Mode]string{
	ModeInfo:   "info",
	ModeError:  "error",
	ModeDebug:  "debug",
	ModeAccess: "access",
}
