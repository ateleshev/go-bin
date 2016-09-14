package strings

func IsASCII(r rune) bool { // {{{
	if r <= MaxASCII {
		return true
	}

	return false
} // }}}

func IsLatin1(r rune) bool { // {{{
	return r <= MaxLatin1
} // }}}

// --[ 0-9 ]--
func IsNumeric(r rune) bool { // {{{
	return IsLatin1(r) && ('0' <= r && r <= '9')
} // }}}

// --[ A-Za-z ]--
func IsAlphabetic(r rune) bool { // {{{
	return IsLatin1(r) && (('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z'))
} // }}}

// --[ 0-9A-Za-z ]--
func IsAlphanumeric(r rune) bool { // {{{
	return IsNumeric(r) || IsAlphabetic(r)
} // }}}

// --[ 0-9A-Za-z_ ]--
func IsWordCharacter(r rune) bool { // {{{
	return IsAlphanumeric(r) || (r == '_')
} // }}}
