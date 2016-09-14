package strings

import "reflect"
import "unsafe"

/**
 * Converts string to a byte slice without memory allocation.
 *
 * Note: it may break if string and/or slice header will change in the future go versions.
 */
func ToBytes(s string) []byte { // {{{
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
} // }}}

func PtrToBytes(s *string) []byte { // {{{
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
} // }}}
