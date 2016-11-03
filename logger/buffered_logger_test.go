package logger

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestBufferedLogger(t *testing.T) {
	var err error
	dir, err := ioutil.TempDir("", "test-buffered-logger_")
	if err != nil {
		t.Fatalf("Cannot create tmp dir: %v", err)
	}
	defer os.RemoveAll(dir) // clean up
	// t.Log(dir)

	l := NewBufferedLogger(NewSeparateFilesLogger("test", dir, ModeAll), 10)
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

	for mode := range ModeNames {
		filename := DefaultFilenameForModeBuilder(dir, l.Name(), mode)

		b, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatalf("Cannot read tmp file: %v", err)
		}

		n := bytes.Count(b, []byte("\n"))
		if n != 2 {
			t.Fatalf("Found incorrect number of lines: %d != 2", n)
		}
	}
}
