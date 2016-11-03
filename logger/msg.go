package logger

type msgChan chan *msg

func newMsg(v ...interface{}) *msg { // {{{
	m := msgPoolGet()
	m.values = v
	return m
} // }}}

func newFmtMsg(f string, v ...interface{}) *msg { // {{{
	m := msgPoolGet()
	m.format = f
	m.values = v
	return m
} // }}}

type msg struct {
	format string
	values []interface{}
}

func (this *msg) Formatted() bool { // {{{
	return this.format != ""
} // }}}

func (this *msg) Reset() { // {{{
	this.format = ""
	this.values = this.values[:0]
} // }}}

func (this *msg) Release() { // {{{
	msgPoolPut(this)
} // }}}
