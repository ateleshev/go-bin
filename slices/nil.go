// Package slices provides primitives for working with slices
package slices

import "reflect"

// Check value is nil
func isNil(v interface{}) bool {
	return v == nil || reflect.ValueOf(v).IsNil()
}

// Move the nil to down of slice
func NilDown(data []interface{}) (l int) {
	// Find non nil at the end of
	i := len(data)
	for i > 0 {
		if !isNil(data[i-1]) {
			break
		}
		i--
	}

	l = i // new len of slice
	i--
	for i >= 0 {
		if isNil(data[i]) {
			data[i] = data[l-1]
			l--
		}
		i--
	}

	return
}
