package bytes

import "unsafe"

/**
 * Converts byte slice to a string without memory allocation.
 *
 * Details: https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ
 * Note: it may break if string and/or slice header will change in the future go versions.
 */
func ToString(b []byte) string { // {{{
	return *(*string)(unsafe.Pointer(&b))
} // }}}
