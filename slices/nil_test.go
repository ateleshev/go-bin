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

func TestSortNilInEnd(t *testing.T) {
	cases := []struct {
		in, want []interface{}
		index    int
	}{
		{[]interface{}{nil, nil, nil}, []interface{}{nil, nil, nil}, 0},
		{[]interface{}{&time.Time{}, nil, nil}, []interface{}{&time.Time{}, nil, nil}, 1},
		{[]interface{}{nil, &time.Time{}, nil}, []interface{}{&time.Time{}, nil, nil}, 1},
		{[]interface{}{nil, nil, &time.Time{}}, []interface{}{&time.Time{}, nil, nil}, 1},
		{[]interface{}{&time.Time{}, &time.Time{}, nil}, []interface{}{&time.Time{}, &time.Time{}, nil}, 2},
		{[]interface{}{&time.Time{}, &time.Time{}}, []interface{}{&time.Time{}, &time.Time{}}, 2},
	}
	for _, c := range cases {
		input := fmt.Sprintf("%q", c.in)
		i, err := SortNilInEnd(c.in)

		if err != nil {
			t.Fatalf("SortNilInEnd(%s) Error: %v", input, err)
		}

		if !reflect.DeepEqual(c.in, c.want) || i != c.index {
			t.Errorf("SortNilInEnd(%s|%d) != %q|%d", input, i, c.want, c.index)
		}
	}
}

/**
 * go test -v -run=^$ -bench=.
 */

func BenchmarkSortNilInEnd(b *testing.B) {
	s := []interface{}{nil, &time.Time{}, &time.Time{}, nil, &time.Time{}, nil, nil}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SortNilInEnd(s)
	}
}

func ExampleSortNilInEnd() {
	s := []interface{}{nil, nil, &time.Time{}, nil, nil}
	n, err := SortNilInEnd(s)
	if err == nil {
		fmt.Println(s)
		fmt.Println(n)
	}
	// Output:
	// [0001-01-01 00:00:00 +0000 UTC <nil> <nil> <nil> <nil>]
	// 1
}
