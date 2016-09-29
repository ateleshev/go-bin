package queue

type Creator func() interface{}

type Result interface {
	Id() int
	Create() interface{}

	Err() error
	Value() interface{}

	Init(string)
	Bind(interface{}, error)

	Metric() Metric

	Reset()
	Release()
}

func NewResult(id int, creator Creator) Result { // {{{
	return &result{
		id:      id,
		creator: creator,
	}
} // }}}

type result struct {
	id      int
	creator Creator

	err   error
	value interface{}

	metric Metric
}

func (this *result) Id() int { // {{{
	return this.id
} // }}}

func (this *result) Create() interface{} { // {{{
	return this.creator()
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

func (this *result) Init(executor string) { // {{{
	this.metric = NewMetric(executor)
} // }}}

func (this *result) Bind(value interface{}, err error) { // {{{
	this.err = err
	this.value = value

	if this.metric != nil {
		this.metric.Calculate(value)
	}
} // }}}

func (this *result) Reset() { // {{{
	this.id = 0
	this.creator = nil
	this.err = nil
	this.value = nil
	this.metric = nil
} // }}}

func (this *result) Release() { // {{{
	resultPoolPut(this)
} // }}}
