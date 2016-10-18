package logger

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileLogger(t *testing.T) {
	var err error
	dir, err := ioutil.TempDir("", "test-file-logger_")
	if err != nil {
		t.Fatalf("Cannot create tmp dir: %v", err)
	}
	defer os.RemoveAll(dir) // clean up
	// t.Log(dir)

	l := NewFileLogger("test", dir, ModeAll)
	if err = l.Open(); err != nil {
		t.Fatalf("Cannot open logger: %v", err)
	}
	defer l.Close()

	l.Access("Access info")
	l.Accessf("Access to: %s", l.Name())

	l.Error("Error info")
	l.Errorf("Error in: %s", l.Name())

	l.Debug("Debug info")
	l.Debugf("Debug in: %s", l.Name())

	filename := DefaultFilenameBuilder(dir, l.Name())
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Cannot read tmp file: %v", err)
	}

	n := bytes.Count(b, []byte("\n"))
	if n != 6 {
		t.Fatalf("Found incorrect number of lines: %d != 6", n)
	}
}
