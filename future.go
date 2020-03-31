package parallel

import "sync"

type Future struct {
	completeWG sync.WaitGroup
	result     interface{}
	err        error
}

func (fut *Future) Get() (interface{}, error) {
	fut.completeWG.Wait()
	return fut.result, fut.err
}

func NewFuture(fn func() (interface{}, error)) *Future {
	fut := &Future{}
	fut.completeWG.Add(1)
	go func() {
		fut.result, fut.err = fn()
		fut.completeWG.Done()
	}()
	return fut
}
