package byteutil

import "unicode/utf8"

var (
	lowUppDiff = byte('a' - 'A')
)

func ToLower(b byte) byte { // {{{
	if b < utf8.RuneSelf {
		if 'A' <= b && b <= 'Z' {
			return byte(b + lowUppDiff)
		}
	}
	return b
} // }}}
