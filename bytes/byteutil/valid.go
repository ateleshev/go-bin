package byteutil

import "unicode"

func IsASCII(v byte) bool { // {{{
	return v <= unicode.MaxASCII
} // }}}

// --[ 0-9 ]--
func IsNumeric(v byte) bool { // {{{
	return '0' <= v && v <= '9'
} // }}}

// --[ A-Za-z ]--
func IsAlphabetic(v byte) bool { // {{{
	return ('A' <= v && v <= 'Z') || ('a' <= v && v <= 'z')
} // }}}

// --[ 0-9A-Za-z ]--
func IsAlphanumeric(v byte) bool { // {{{
	return IsNumeric(v) || IsAlphabetic(v)
} // }}}

// --[ 0-9A-Za-z_ ]--
func IsWordCharacter(v byte) bool { // {{{
	return IsAlphanumeric(v) || (v == '_')
} // }}}
