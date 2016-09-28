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

	Reset()
	Release()
}

func NewResults(id int, creator Creator) Results { // {{{
	return &results{
		id:      id,
		creator: creator,
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

	if this.metrics != nil {
		this.metrics.Calculate(value)
	}
} // }}}

func (this *results) Reset() { // {{{
	this.id = 0
	this.creator = nil
	this.err = nil
	this.value = nil
	this.metrics = nil
} // }}}

func (this *results) Release() { // {{{
	resultsPoolPut(this)
} // }}}
