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

	l.Info("Info")
	l.Infof("Info logger name: %s", l.Name())

	l.Error("Error")
	l.Errorf("Error logger name: %s", l.Name())

	l.Debug("Debug")
	l.Debugf("Debug logger name: %s", l.Name())

	l.Access("Access")
	l.Accessf("Access logger name: %s", l.Name())

	if err := l.Close(); err != nil {
		t.Fatal(err)
	}

	filename := DefaultFilenameBuilder(dir, l.Name())
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Cannot read tmp file: %v", err)
	}

	n := bytes.Count(b, []byte("\n"))
	if n != 8 {
		t.Fatalf("Found incorrect number of lines: %d != 6", n)
	}
}
