package queue

import "time"
import "unsafe"

type Metric interface {
	Executor() string
	BeginTime() time.Time
	FinishTime() time.Time

	Init()
	Calculate(interface{})

	ElapsedTime() time.Duration
	MemoryUsage() uintptr

	Reset()
	Release()
}

func NewMetric(executor string) Metric { // {{{
	m := metricPoolGet()
	m.executor = executor
	m.Init()
	return m
} // }}}

type metric struct {
	executor    string
	beginTime   time.Time
	finishTime  time.Time
	memoryUsage uintptr
}

func (this *metric) Reset() { // {{{
	this.executor = ""
	this.beginTime = time.Time{}
	this.finishTime = time.Time{}
	this.memoryUsage = 0
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

func (this *metric) Init() { // {{{
	m.beginTime = time.Now()
} // }}}

func (this *metric) Calculate(v interface{}) { // {{{
	this.finishTime = time.Now()
	this.memoryUsage = unsafe.Sizeof(v)
} // }}}

func (this *metric) ElapsedTime() time.Duration { // {{{
	return this.finishTime.Sub(this.beginTime)
} // }}}

func (this *metric) MemoryUsage() uintptr { // {{{
	return this.memoryUsage
} // }}}
