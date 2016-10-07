// Package slices provides primitives for working with slices
package slices

import "reflect"

func clear(v reflect.Value) interface{} {
	n := sortNilInEnd(v)
	return v.Slice(0, n).Interface()
}

// in - inputed slice
// out - outputed slice
func Clear(in interface{}) (out interface{}, err error) {
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Slice {
		err = ErrItIsNotSlice
		return
	}

	out = clear(v)
	return
}
