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
		input := fmt.Sprintf("%q", c.in)
		res, err := Clear(c.in)

		if err != nil {
			t.Fatalf("Clear(%s) Error: %v", input, err)
		}

		if !reflect.DeepEqual(res.([]interface{}), c.want) {
			t.Errorf("Clear(%s) != %q", input, c.want)
		}
	}
}

/**
 * go test -v -run=^$ -benchmem -bench=.
 */

func BenchmarkClear(b *testing.B) {
	var err error
	s := []interface{}{nil, &time.Time{}, &time.Time{}, nil, &time.Time{}, nil, nil}
	input := fmt.Sprintf("%q", s)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = Clear(s)
		if err != nil {
			b.Fatalf("Clear(%s) Error: %v", input, err)
		}
	}
}

func ExampleClear() {
	s := []interface{}{&time.Time{}, nil, nil}
	r, err := Clear(s)
	if err == nil {
		fmt.Println(r.([]interface{}))
	}
	// Output:
	// [0001-01-01 00:00:00 +0000 UTC]
}
