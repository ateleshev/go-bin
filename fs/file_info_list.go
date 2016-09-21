package fs

/**
 * fs - file system
 */

import "os"
import "sync"

func NewFileInfoList() FileInfoList { // {{{
	return &fileInfoList{
		mu:  new(sync.Mutex),
		val: make([]os.FileInfo, 0),
	}
} // }}}

type FileInfoList interface {
	Append(os.FileInfo)
	Value() []os.FileInfo
	Count() int
	Has() bool
	First() os.FileInfo
	Last() os.FileInfo
	Reset()
}

type fileInfoList struct {
	mu  *sync.Mutex
	val []os.FileInfo
}

func (this *fileInfoList) Append(v os.FileInfo) { // {{{
	this.mu.Lock()
	defer this.mu.Unlock()
	this.val = append(this.val, v)
} // }}}

func (this *fileInfoList) Value() []os.FileInfo { // {{{
	return this.val
} // }}}

func (this *fileInfoList) Count() int { // {{{
	return len(this.val)
} // }}}

func (this *fileInfoList) Has() bool { // {{{
	return this.Count() > 0
} // }}}

func (this *fileInfoList) First() os.FileInfo { // {{{
	if this.Has() {
		return this.val[0]
	}
	return nil
} // }}}

func (this *fileInfoList) Last() os.FileInfo { // {{{
	if this.Has() {
		return this.val[this.Count()-1]
	}
	return nil
} // }}}

func (this *fileInfoList) Reset() { // {{{
	this.val = this.val[:0]
} // }}}
