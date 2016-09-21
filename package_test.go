package bin

/**
 * Profiling of test coverage
 * #> go test -coverprofile=cover.out
 *
 * Analyze results of test coverage
 * #> go tool cover -func=cover.out
 * #> go tool cover -html=cover.out
 */

import "testing"
import "github.com/ateleshev/go-bin/bytes"
import "github.com/ateleshev/go-bin/strings"

func TestBool(t *testing.T) { // {{{
	if (true != false) != true {
		t.Error("Something is wrong.")
	}
} // }}}

func TestImports(t *testing.T) { // {{{
	s := "Test go-bin package"
	b := []byte("Test go-bin package")

	if !strings.Equal(s, bytes.ToString(b)) {
		t.Error("Conversion of bytes to string does not work")
	}

	if !bytes.Equal(b, strings.ToBytes(s)) {
		t.Error("Conversion of string to bytes does not work")
	}
} // }}}
