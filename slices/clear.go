// Package slices provides primitives for working with slices
package slices

// Clear from nil
func Clear(data []interface{}) []interface{} {
	l := NilDown(data)
	return data[:l]
}
