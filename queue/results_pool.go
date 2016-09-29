// go-bin: github.com/ateleshev/go-bin
package queue

import "sync"

var resultPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &result{}
	},
} // }}}

func resultPoolGet() *result { // {{{
	if instance := resultPool.Get(); instance != nil {
		return instance.(*result)
	}

	return resultPool.New().(*result)
} // }}}

func resultPoolPut(instance *result) { // {{{
	instance.Reset()
	resultPool.Put(instance)
} // }}}
