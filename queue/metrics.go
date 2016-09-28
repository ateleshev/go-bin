package queue

import "time"
import "unsafe"

type Metrics interface {
	Executor() string
	BeginTime() time.Time
	FinishTime() time.Time

	Calculate(interface{})

	ElapsedTime() time.Duration
	MemoryUsage() uintptr
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
	memoryUsage uintptr
}

func (this *metrics) Reset() { // {{{
	this.executor = ""
	this.beginTime = time.Time{}
	this.finishTime = time.Time{}
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
	this.memoryUsage = unsafe.Sizeof(v)
} // }}}

func (this *metrics) ElapsedTime() time.Duration { // {{{
	return this.finishTime.Sub(this.beginTime)
} // }}}

func (this *metrics) MemoryUsage() uintptr { // {{{
	return this.memoryUsage
} // }}}
