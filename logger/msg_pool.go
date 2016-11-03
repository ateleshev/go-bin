// go-bin: github.com/ateleshev/go-bin
package logger

import (
	"sync"
)

var msgPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &msg{}
	},
} // }}}

func msgPoolGet() *msg { // {{{
	if instance := msgPool.Get(); instance != nil {
		return instance.(*msg)
	}

	return msgPool.New().(*msg)
} // }}}

func msgPoolPut(instance *msg) { // {{{
	instance.Reset()
	msgPool.Put(instance)
} // }}}
