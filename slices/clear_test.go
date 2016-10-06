// Package slices provides primitives for working with slices
package slices

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

/**
 * go test -v -race
 */

func TestClear(t *testing.T) {
	cases := []struct {
		in, want []interface{}
	}{
		{[]interface{}{nil, nil, nil}, []interface{}{}},
		{[]interface{}{&time.Time{}, nil, nil}, []interface{}{&time.Time{}}},
		{[]interface{}{nil, &time.Time{}, nil}, []interface{}{&time.Time{}}},
		{[]interface{}{nil, nil, &time.Time{}}, []interface{}{&time.Time{}}},
		{[]interface{}{&time.Time{}, &time.Time{}, nil}, []interface{}{&time.Time{}, &time.Time{}}},
		{[]interface{}{&time.Time{}, &time.Time{}}, []interface{}{&time.Time{}, &time.Time{}}},
	}
	for _, c := range cases {
		in := fmt.Sprintf("%q", c.in)
		if !reflect.DeepEqual(Clear(c.in), c.want) {
			t.Errorf("Clear(%s) != %q", in, c.want)
		}
	}
}

/**
 * go test -v -run=^$ -benchmem -bench=.
 */

func BenchmarkClear(b *testing.B) {
	sl := []interface{}{nil, &time.Time{}, &time.Time{}, nil, &time.Time{}, nil, nil}
	for i := 0; i < b.N; i++ {
		Clear(sl)
	}
}
