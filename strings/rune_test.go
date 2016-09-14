package strings_test

import "testing"
import "github.com/ateleshev/go-bin/strings"

func TestIsLatin1(t *testing.T) { // {{{
	if !strings.IsLatin1('0') {
		t.Fatal("'0' - is Latin1")
	}
	if !strings.IsLatin1('3') {
		t.Fatal("'3' - is Latin1")
	}
	if !strings.IsLatin1('6') {
		t.Fatal("'6' - is Latin1")
	}
	if !strings.IsLatin1('9') {
		t.Fatal("'9' - is Latin1")
	}
	if !strings.IsLatin1('A') {
		t.Fatal("'A' - is Latin1")
	}
	if !strings.IsLatin1('T') {
		t.Fatal("'T' - is Latin1")
	}
	if !strings.IsLatin1('Z') {
		t.Fatal("'Z' - is Latin1")
	}
	if !strings.IsLatin1('a') {
		t.Fatal("'a' - is Latin1")
	}
	if !strings.IsLatin1('k') {
		t.Fatal("'k' - is Latin1")
	}
	if !strings.IsLatin1('z') {
		t.Fatal("'z' - is Latin1")
	}
	if strings.IsLatin1('ы') {
		t.Log("'ы' - is not a Latin1")
	}
	if strings.IsLatin1('ї') {
		t.Log("'ї' - is not a Latin1")
	}
	if strings.IsLatin1('і') {
		t.Log("'і' - is not a Latin1")
	}
} // }}}
