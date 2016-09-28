// go-bin: github.com/ateleshev/go-bin
package queue

import (
	"sync"
)

var metricsPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &metrics{}
	},
} // }}}

func metricsPoolGet() *metrics { // {{{
	if instance := metricsPool.Get(); instance != nil {
		return instance.(*metrics)
	}

	return metricsPool.New().(*metrics)
} // }}}

func metricsPoolPut(instance *metrics) { // {{{
	instance.Reset()
	metricsPool.Put(instance)
} // }}}
