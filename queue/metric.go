package queue

import "time"
import "github.com/ateleshev/go-bin/cmp"

type Metric interface {
	Name() string
	BeginTime() time.Time
	FinishTime() time.Time
	ElapsedTime() time.Duration

	Begin()
	Finish()

	cmp.ResetReleaser
}

func NewMetric(name string) Metric { // {{{
	m := metricPoolGet()
	m.name = name
	return m
} // }}}

type metric struct {
	name       string
	beginTime  time.Time
	finishTime time.Time
}

func (this *metric) Reset() { // {{{
	this.name = ""
	//	this.beginTime = time.Time{}
	//	this.finishTime = time.Time{}
} // }}}

func (this *metric) Release() { // {{{
	metricPoolPut(this)
} // }}}

func (this *metric) Name() string { // {{{
	return this.name
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
