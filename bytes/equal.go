package bytes

func Equal(a, b []byte) bool { // {{{
	var i, l int
	l = len(a)

	if l != len(b) {
		return false
	}

	for i < l {
		if a[i] != b[i] {
			return false
		}
		i++
	}

	return true
} // }}}
