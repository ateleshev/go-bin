package es_test

import "errors"
import "reflect"
import "testing"
import "github.com/ateleshev/go-bin/es"

var (
	esList = es.NewErrorList()
	esData = []error{
		errors.New("First"),
		errors.New("Next"),
		errors.New("Last"),
	}
)

func TestErrorList(t *testing.T) {
	if _, ok := esList.(es.ErrorList); !ok {
		t.Fatal("Was created incorrect instance")
	}

	if esList.Has() {
		t.Fatal("Has - does not work")
	}

	if !reflect.DeepEqual([]error{}, esList.Value()) {
		t.Fatal("Value - does not work")
	}

	if esList.First() != nil {
		t.Fatal("First - does not work")
	}

	if esList.Last() != nil {
		t.Fatal("Last - does not work")
	}

	for _, e := range esData {
		esList.Append(e)
	}

	if !esList.Has() {
		t.Fatal("Append - does not work")
	}

	if esList.Count() != 3 {
		t.Fatalf("Count - does not work: 3 != %d", esList.Count())
	}

	if !reflect.DeepEqual(esData, esList.Value()) {
		t.Fatal("Value - does not work")
	}

	if !reflect.DeepEqual(esData[0], esList.First()) {
		t.Fatal("First - does not work")
	}

	if !reflect.DeepEqual(esData[2], esList.Last()) {
		t.Fatal("Last - does not work")
	}

	esList.Reset()

	if esList.Has() {
		t.Fatal("Reset - does not work")
	}

	if esList.Count() != 0 {
		t.Fatalf("Count - does not work: 0 != %d", esList.Count())
	}
}
