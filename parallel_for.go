package parallel

import (
	"bytes"
	"errors"
	"sync"
)

type parallelFor struct {
	wg        sync.WaitGroup
	errorList []error
	errorChan chan error
	errLoopWg sync.WaitGroup
}

func (pf *parallelFor) errorCollectionLoop() {
	for {
		err := <-pf.errorChan
		if err == nil {
			break
		}
		pf.errorList = append(pf.errorList, err)
	}
	pf.errLoopWg.Done()
}

func (pf *parallelFor) wait() error {
	pf.wg.Wait()
	pf.errorChan <- nil
	pf.errLoopWg.Wait()
	if len(pf.errorList) == 0 {
		return nil
	}
	errStr := bytes.NewBuffer(nil)
	errStr.WriteString("errored in parallel for:")
	for _, e := range pf.errorList {
		errStr.WriteString("\n")
		errStr.WriteString(e.Error())
	}
	return errors.New(errStr.String())
}

// For will execute N goroutines, wait them complete, and gather errors.
//
// The callback `fn` takes two parameters: the offset of goroutine and the total size of goroutines.
func For(n int, fn func(i int, total int) error) error {
	pf := &parallelFor{
		wg:        sync.WaitGroup{},
		errorList: nil,
		errorChan: make(chan error),
		errLoopWg: sync.WaitGroup{},
	}
	pf.wg.Add(n)
	pf.errLoopWg.Add(1)
	go pf.errorCollectionLoop()
	for i := 0; i < n; i++ {
		offset := i
		go func() {
			err := fn(offset, n)
			if err != nil {
				pf.errorChan <- err
			}
			pf.wg.Done()
		}()
	}

	return pf.wait()
}
