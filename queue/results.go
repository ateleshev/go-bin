package queue

type Creator func() interface{}

type Results interface {
	Id() int
	Create() interface{}

	Err() error
	Value() interface{}

	Init(string)
	Bind(interface{}, error)

	Metrics() Metrics
}

func Newresults(id int, creator Creator) results { // {{{
	return &results{
		id: id,
	}
} // }}}

type results struct {
	id      int
	creator Creator

	err   error
	value interface{}

	metrics Metrics
}

func (this *results) Id() int { // {{{
	return this.id
} // }}}

func (this *results) Create() interface{} { // {{{
	return this.creator()
} // }}}

func (this *results) Err() error { // {{{
	return this.err
} // }}}

func (this *results) Value() interface{} { // {{{
	return this.value
} // }}}

func (this *results) Metrics() Metrics { // {{{
	return this.metrics
} // }}}

func (this *results) Init(executor string) { // {{{
	this.metrics = NewMetrics(executor)
} // }}}

func (this *results) Bind(value interface{}, err error) { // {{{
	this.err = err
	this.value = value
	this.metrics.Calculate(value)
} // }}}
