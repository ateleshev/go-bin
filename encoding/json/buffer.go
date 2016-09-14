package json

import "sync"

const (
	numericValueBufferCap = 20
)

var numericValueBufferPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &numericValueBuffer{
			B: make([]byte, 0, numericValueBufferCap),
		}
	},
} // }}}

func getNumericValueBuffer() *numericValueBuffer { // {{{
	if instance := numericValueBufferPool.Get(); instance != nil {
		return instance.(*numericValueBuffer)
	}

	return numericValueBufferPool.New().(*numericValueBuffer)
} // }}}

func putNumericValueBuffer(instance *numericValueBuffer) { // {{{
	instance.Reset()
	numericValueBufferPool.Put(instance)
} // }}}

func newNumericValueBuffer() *numericValueBuffer { // {{{
	return getNumericValueBuffer()
} // }}}

type numericValueBuffer struct {
	B []byte
}

func (this *numericValueBuffer) Reset() { // {{{
	this.B = this.B[:0]
} // }}}

func (this *numericValueBuffer) Release() { // {{{
	putNumericValueBuffer(this)
} // }}}
