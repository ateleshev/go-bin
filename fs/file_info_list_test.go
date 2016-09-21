package fs_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ateleshev/go-bin/fs"
)

func newFakeFileInfo(n string) *fakeFileInfo { // {{{
	return &fakeFileInfo{name: n}
} // }}}

type fakeFileInfo struct {
	name string
}

func (this *fakeFileInfo) Name() string { // {{{
	return this.name
} // }}}

func (this *fakeFileInfo) Size() int64 { // {{{
	return int64(0)
} // }}}

func (this *fakeFileInfo) Mode() os.FileMode { // {{{
	return os.ModePerm
} // }}}

func (this *fakeFileInfo) ModTime() time.Time { // {{{
	return time.Now()
} // }}}

func (this *fakeFileInfo) IsDir() bool { // {{{
	return false
} // }}}

func (this *fakeFileInfo) Sys() interface{} { // {{{
	return this
} // }}}

var (
	fiList = fs.NewFileInfoList()
	fiData = []os.FileInfo{
		newFakeFileInfo("First"),
		newFakeFileInfo("Next"),
		newFakeFileInfo("Last"),
	}
)

func TestHolder(t *testing.T) {
	if _, ok := fiList.(fs.FileInfoList); !ok {
		t.Fatal("Was created incorrect instance")
	}

	if fiList.Has() {
		t.Fatal("Has - does not work")
	}

	if !reflect.DeepEqual([]os.FileInfo{}, fiList.Value()) {
		t.Fatal("Value - does not work")
	}

	if fiList.First() != nil {
		t.Fatal("First - does not work")
	}

	if fiList.Last() != nil {
		t.Fatal("Last - does not work")
	}

	for _, fi := range fiData {
		fiList.Append(fi)
	}

	if !fiList.Has() {
		t.Fatal("Append - does not work")
	}

	if fiList.Count() != 3 {
		t.Fatalf("Count - does not work: 3 != %d", fiList.Count())
	}

	if !reflect.DeepEqual(fiData, fiList.Value()) {
		t.Fatal("Value - does not work")
	}

	if !reflect.DeepEqual(fiData[0], fiList.First()) {
		t.Fatal("First - does not work")
	}

	if !reflect.DeepEqual(fiData[2], fiList.Last()) {
		t.Fatal("Last - does not work")
	}

	fiList.Reset()

	if fiList.Has() {
		t.Fatal("Reset - does not work")
	}

	if fiList.Count() != 0 {
		t.Fatalf("Count - does not work: 0 != %d", fiList.Count())
	}
}
