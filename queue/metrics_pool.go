// go-bin: github.com/ateleshev/go-bin
package queue

import "sync"

var metricPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &metric{}
	},
} // }}}

func metricPoolGet() *metric { // {{{
	if instance := metricPool.Get(); instance != nil {
		return instance.(*metric)
	}

	return metricPool.New().(*metric)
} // }}}

func metricPoolPut(instance *metric) { // {{{
	instance.Reset()
	metricPool.Put(instance)
} // }}}
