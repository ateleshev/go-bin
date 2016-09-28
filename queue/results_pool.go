// go-bin: github.com/ateleshev/go-bin
package queue

import (
	"sync"
)

var resultsPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &results{}
	},
} // }}}

func resultsPoolGet() *results { // {{{
	if instance := resultsPool.Get(); instance != nil {
		return instance.(*results)
	}

	return resultsPool.New().(*results)
} // }}}

func resultsPoolPut(instance *results) { // {{{
	instance.Reset()
	resultsPool.Put(instance)
} // }}}
