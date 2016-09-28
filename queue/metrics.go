package queue

import "time"
import "encoding/binary"

type Metrics interface {
	Executor() string
	BeginTime() time.Time
	FinishTime() time.Time

	Calculate(interface{})

	ElapsedTime() time.Duration
	MemoryUsage() int
}

func NewMetrics(executor string) Metrics { // {{{
	m := metricsPoolGet()
	m.executor = executor
	m.beginTime = time.Now()
	return m
} // }}}

type metrics struct {
	executor    string
	beginTime   time.Time
	finishTime  time.Time
	memoryUsage int
}

func (this *metrics) Reset() { // {{{
	this.executor = ""
	this.beginTime = (time.Time)(nil)
	this.finishTime = (time.Time)(nil)
	this.memoryUsage = 0
} // }}}

func (this *metrics) Release() { // {{{
	metricsPoolPut(this)
} // }}}

func (this *metrics) Executor() string { // {{{
	return this.executor
} // }}}

func (this *metrics) BeginTime() time.Time { // {{{
	return this.beginTime
} // }}}

func (this *metrics) FinishTime() time.Time { // {{{
	return this.finishTime
} // }}}

func (this *metrics) Calculate(v interface{}) { // {{{
	this.finishTime = time.Now()
	this.memoryUsage = binary.Size(v)
} // }}}

func (this *metrics) ElapsedTime() time.Duration { // {{{
	return this.endTime.Sub(this.beginTime)
} // }}}

func (this *metrics) MemoryUsage() int { // {{{
	return this.memoryUsage
} // }}}
