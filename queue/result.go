package queue

import "github.com/ateleshev/go-bin/cmp"

type Result interface {
	Id() int
	Err() error
	Value() interface{}

	Init(string)
	Bind(interface{}, error)
	Create() interface{}
	Metric() Metric

	cmp.ResetReleaser
}

func NewResult(id int, creator cmp.Creator) Result { // {{{
	return &result{
		id:      id,
		creator: creator,
	}
} // }}}

type result struct {
	id      int
	err     error
	value   interface{}
	metric  Metric
	creator cmp.Creator
}

func (this *result) Id() int { // {{{
	return this.id
} // }}}

func (this *result) Create() interface{} { // {{{
	return this.creator.Create()
} // }}}

func (this *result) Err() error { // {{{
	return this.err
} // }}}

func (this *result) Value() interface{} { // {{{
	return this.value
} // }}}

func (this *result) Metric() Metric { // {{{
	return this.metric
} // }}}

func (this *result) Init(name string) { // {{{
	this.metric = NewMetric(name)
	this.metric.Begin()
} // }}}

func (this *result) Bind(value interface{}, err error) { // {{{
	this.err = err
	this.value = value

	if this.metric != nil {
		this.metric.Finish()
	}
} // }}}

func (this *result) Reset() { // {{{
	this.id = 0
	this.err = nil
	this.value = nil
	this.metric = nil
	this.creator = nil
} // }}}

func (this *result) Release() { // {{{
	resultPoolPut(this)
} // }}}
