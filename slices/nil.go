// Package slices provides primitives for working with slices
package slices

import "reflect"

// Move the nil to down of slice
func sortNilInEnd(slice reflect.Value) (n int) {
	// Find non nil at the end of
	i := slice.Len()
	for i > 0 {
		if !slice.Index(i - 1).IsNil() {
			break
		}
		i--
	}

	n = i // new len of slice
	i--
	for i >= 0 {
		current := slice.Index(i)
		if current.IsNil() {
			last := slice.Index(n - 1)
			current.Set(last)
			last.Set(reflect.Zero(current.Type()))
			n--
		}
		i--
	}

	return
}

// n - index of first element of nil
func SortNilInEnd(s interface{}) (n int, err error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		err = ErrItIsNotSlice
		return
	}

	n = sortNilInEnd(v)
	return
}
