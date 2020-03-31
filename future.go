package parallel

import (
	"go.uber.org/atomic"
	"sync"
)

type Future struct {
	completeWG sync.WaitGroup
	isComplete *atomic.Bool
	result     interface{}
	err        error
}

func (fut *Future) Get() (interface{}, error) {
	fut.completeWG.Wait()
	return fut.result, fut.err
}

func (fut *Future) IsComplete() bool {
	return fut.isComplete.Load()
}

func NewFuture(fn func() (interface{}, error)) *Future {
	fut := &Future{
		isComplete: atomic.NewBool(false),
	}
	fut.completeWG.Add(1)
	go func() {
		fut.result, fut.err = fn()
		fut.isComplete.Store(true)
		fut.completeWG.Done()
	}()
	return fut
}
