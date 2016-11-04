package byteutil

import (
	"fmt"
)

type Size float64

func (b Size) String() string { // {{{
	switch {
	case b >= YB:
		return fmt.Sprintf("%.3fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.3fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.3fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.3fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.3fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.3fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.3fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.3fKB", b/KB)
	}

	return fmt.Sprintf("%.3fB", b)
} // }}}

const (
	_       = iota // ignore first value by assigning to blank identifier
	KB Size = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
