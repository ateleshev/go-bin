package queue

import "time"

type Metric interface {
	Executor() string
	BeginTime() time.Time
	FinishTime() time.Time
	ElapsedTime() time.Duration

	Begin()
	Finish()

	Reset()
	Release()
}

func NewMetric(executor string) Metric { // {{{
	m := metricPoolGet()
	m.executor = executor
	return m
} // }}}

type metric struct {
	executor   string
	beginTime  time.Time
	finishTime time.Time
}

func (this *metric) Reset() { // {{{
	this.executor = ""
	this.beginTime = time.Time{}
	this.finishTime = time.Time{}
} // }}}

func (this *metric) Release() { // {{{
	metricPoolPut(this)
} // }}}

func (this *metric) Executor() string { // {{{
	return this.executor
} // }}}

func (this *metric) BeginTime() time.Time { // {{{
	return this.beginTime
} // }}}

func (this *metric) FinishTime() time.Time { // {{{
	return this.finishTime
} // }}}

func (this *metric) ElapsedTime() time.Duration { // {{{
	return this.finishTime.Sub(this.beginTime)
} // }}}

func (this *metric) Begin() { // {{{
	this.beginTime = time.Now()
} // }}}

func (this *metric) Finish() { // {{{
	this.finishTime = time.Now()
} // }}}
