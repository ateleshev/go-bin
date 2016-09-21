package holder

import "sync"

func New() Holder { // {{{
	return &holder{
		mu:  new(sync.Mutex),
		val: make([]error, 0),
	}
} // }}}

type Holder interface {
	Append(error)
	Value() []error
	Count() int
	Has() bool
	First() error
	Last() error
	Reset()
}

type holder struct {
	mu  *sync.Mutex
	val []error
}

func (this *holder) Append(e error) { // {{{
	this.mu.Lock()
	defer this.mu.Unlock()
	this.val = append(this.val, e)
} // }}}

func (this *holder) Value() []error { // {{{
	return this.val
} // }}}

func (this *holder) Count() int { // {{{
	return len(this.val)
} // }}}

func (this *holder) Has() bool { // {{{
	return this.Count() > 0
} // }}}

func (this *holder) First() error { // {{{
	if this.Has() {
		return this.val[0]
	}
	return nil
} // }}}

func (this *holder) Last() error { // {{{
	if this.Has() {
		return this.val[this.Count()-1]
	}
	return nil
} // }}}

func (this *holder) Reset() { // {{{
	this.val = this.val[:0]
} // }}}
