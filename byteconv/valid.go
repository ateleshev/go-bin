package byteconv

import "unicode/utf8"

// --[ 0-9 ]--
func IsNumeric(v byte) bool { // {{{
	return v < utf8.RuneSelf && ('0' <= v && v <= '9')
} // }}}

// --[ A-Za-z ]--
func IsAlphabetic(v byte) bool { // {{{
	return v < utf8.RuneSelf && (('A' <= v && v <= 'Z') || ('a' <= v && v <= 'z'))
} // }}}

// --[ 0-9A-Za-z ]--
func IsAlphanumeric(v byte) bool { // {{{
	return IsNumeric(v) || IsAlphabetic(v)
} // }}}

// --[ 0-9A-Za-z_ ]--
func IsWordCharacter(v byte) bool { // {{{
	return IsAlphanumeric(v) || (v == '_')
} // }}}
