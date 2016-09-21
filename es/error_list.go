package es

/**
 * es - errors
 */

import "sync"

func NewErrorList() ErrorList { // {{{
	return &errorList{
		mu:  new(sync.Mutex),
		val: make([]error, 0),
	}
} // }}}

type ErrorList interface {
	Append(error)
	Value() []error
	Count() int
	Has() bool
	First() error
	Last() error
	Reset()
}

type errorList struct {
	mu  *sync.Mutex
	val []error
}

func (this *errorList) Append(v error) { // {{{
	this.mu.Lock()
	defer this.mu.Unlock()
	this.val = append(this.val, v)
} // }}}

func (this *errorList) Value() []error { // {{{
	return this.val
} // }}}

func (this *errorList) Count() int { // {{{
	return len(this.val)
} // }}}

func (this *errorList) Has() bool { // {{{
	return this.Count() > 0
} // }}}

func (this *errorList) First() error { // {{{
	if this.Has() {
		return this.val[0]
	}
	return nil
} // }}}

func (this *errorList) Last() error { // {{{
	if this.Has() {
		return this.val[this.Count()-1]
	}
	return nil
} // }}}

func (this *errorList) Reset() { // {{{
	this.val = this.val[:0]
} // }}}
