package holder_test

import "errors"
import "reflect"
import "testing"
import "github.com/ateleshev/go-bin/errors/holder"

var (
	h  = holder.New()
	es = []error{
		errors.New("First"),
		errors.New("Next"),
		errors.New("Last"),
	}
)

func TestHolder(t *testing.T) {
	if _, ok := h.(holder.Holder); !ok {
		t.Fatal("Was created incorrect instance")
	}

	if h.Has() {
		t.Fatal("Has - does not work")
	}

	if !reflect.DeepEqual([]error{}, h.Value()) {
		t.Fatal("Value - does not work")
	}

	if h.First() != nil {
		t.Fatal("First - does not work")
	}

	if h.Last() != nil {
		t.Fatal("Last - does not work")
	}

	for _, e := range es {
		h.Append(e)
	}

	if !h.Has() {
		t.Fatal("Append - does not work")
	}

	if h.Count() != 3 {
		t.Fatalf("Count - does not work: 3 != %d", h.Count())
	}

	if !reflect.DeepEqual(es, h.Value()) {
		t.Fatal("Value - does not work")
	}

	if !reflect.DeepEqual(es[0], h.First()) {
		t.Fatal("First - does not work")
	}

	if !reflect.DeepEqual(es[2], h.Last()) {
		t.Fatal("Last - does not work")
	}

	h.Reset()

	if h.Has() {
		t.Fatal("Reset - does not work")
	}

	if h.Count() != 0 {
		t.Fatalf("Count - does not work: 0 != %d", h.Count())
	}
}
