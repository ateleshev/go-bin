package queue

type Results interface {
	Id() int
	Err() error
	Value() interface{}
	Metrics() Metrics

	Init(string)
	Bind(interface{}, error)
}

func NewResult(id int) Result { // {{{
	return &result{
		id: id,
	}
} // }}}

type result struct {
	id      int
	err     error
	value   interface{}
	metrics Metrics
}

func (this *result) Id() int { // {{{
	return this.id
} // }}}

func (this *result) Err() error { // {{{
	return this.err
} // }}}

func (this *result) Value() interface{} { // {{{
	return this.value
} // }}}

func (this *result) Metrics() Metrics { // {{{
	return this.metrics
} // }}}

func (this *result) Init(executor string) { // {{{
	this.metrics = NewMetrics(executor)
} // }}}

func (this *result) Bind(value interface{}, err error) { // {{{
	this.err = err
	this.value = value
	this.metrics.Calculate(value)
} // }}}
