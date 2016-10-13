package queue

import "time"
import "runtime"
import "sync/atomic"
import "github.com/ateleshev/go-bin/cmp"

type Metric interface {
	Name() string
	Number() int64
	BeginTime() time.Time
	FinishTime() time.Time
	ElapsedTime() time.Duration

	Increase() int64
	Decrease() int64

	Begin()
	Finish()

	/**
	 * Details: https://golang.org/pkg/runtime/#MemStats
	 */
	// runtime.MemStats: Alloc, Sys, TotalAlloc
	MemUsage() (uint64, uint64, uint64)
	// runtime.MemStats: HeapAlloc, HeapSys, HeapObjects
	HeapUsage() (uint64, uint64, uint64)
	// runtime.MemStats: StackInuse, StackSys, OtherSys
	StackUsage() (uint64, uint64, uint64)

	cmp.ResetReleaser
}

func NewMetric(name string) Metric { // {{{
	m := metricPoolGet()
	m.name = name
	return m
} // }}}

type metric struct {
	name        string
	number      int64
	beginTime   time.Time
	finishTime  time.Time
	elapsedTime time.Duration
	memStats    runtime.MemStats
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

func (this *metric) Number() int64 { // {{{
	return this.number
} // }}}

func (this *metric) BeginTime() time.Time { // {{{
	return this.beginTime
} // }}}

func (this *metric) FinishTime() time.Time { // {{{
	return this.finishTime
} // }}}

func (this *metric) ElapsedTime() time.Duration { // {{{
	return this.elapsedTime
} // }}}

func (this *metric) Increase() int64 { // {{{
	return atomic.AddInt64(&this.number, 1)
} // }}}

func (this *metric) Decrease() int64 { // {{{
	return atomic.AddInt64(&this.number, -1)
} // }}}

func (this *metric) Begin() { // {{{
	this.beginTime = time.Now()
} // }}}

func (this *metric) Finish() { // {{{
	this.finishTime = time.Now()
	this.elapsedTime = this.finishTime.Sub(this.beginTime)
	runtime.ReadMemStats(&this.memStats)
} // }}}

/**
 * General statistics.
 *
 * Alloc      uint64 // bytes allocated and not yet freed
 * Sys        uint64 // bytes obtained from system (sum of XxxSys below)
 * TotalAlloc uint64 // bytes allocated (even if freed)
 */
func (this *metric) MemUsage() (uint64, uint64, uint64) { // {{{
	return this.memStats.Alloc, this.memStats.Sys, this.memStats.TotalAlloc
} // }}}

/**
 * Main allocation heap statistics.
 *
 * HeapAlloc    uint64 // bytes allocated and not yet freed (same as Alloc above)
 * HeapSys      uint64 // bytes obtained from system
 * HeapObjects  uint64 // total number of allocated objects
 */
func (this *metric) HeapUsage() (uint64, uint64, uint64) { // {{{
	return this.memStats.HeapAlloc, this.memStats.HeapSys, this.memStats.HeapObjects
} // }}}

/**
 * Low-level fixed-size structure allocator statistics.
 *   Inuse is bytes used now.
 *   Sys is bytes obtained from system.
 *
 * StackInuse  uint64 // bytes used by stack allocator
 * StackSys    uint64
 * OtherSys    uint64 // other system allocations
 */
func (this *metric) StackUsage() (uint64, uint64, uint64) { // {{{
	return this.memStats.StackInuse, this.memStats.StackSys, this.memStats.OtherSys
} // }}}
