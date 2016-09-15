package dom

import "sync"

var elmFinderPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &ElmFinder{
			result: make([]byte, ElmFinderResultCap),
		}
	},
} // }}}

func getElmFinder() *ElmFinder { // {{{
	if instance := elmFinderPool.Get(); instance != nil {
		return instance.(*ElmFinder)
	}

	return elmFinderPool.New().(*ElmFinder)
} // }}}

func putElmFinder(instance *ElmFinder) { // {{{
	instance.Reset()
	elmFinderPool.Put(instance)
} // }}}
