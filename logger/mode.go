package logger

type Mode uint8

// Modes of loggers: (Access | Error | Debug)
const (
	// modes quantity
	ModesQnt = 3

	ModeAccess Mode = 1 << iota
	ModeError
	ModeDebug

	ModeAll        Mode = ModeAccess | ModeError | ModeDebug
	ModeErrorDebug Mode = ModeError | ModeDebug
)

var ModeName map[Mode]string = map[Mode]string{
	ModeAccess: "access",
	ModeError:  "error",
	ModeDebug:  "debug",
}
