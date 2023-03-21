package storage

import (
	"context"
	"sync"
	"sync/atomic"
)

// Result represents the Size function result
type Result struct {
	// Total Size of File objects
	Size int64
	// Count is a count of File objects processed
	Count int64
}

type DirSizer interface {
	// Size calculate a size of given Dir, receive a ctx and the root Dir instance
	// will return Result or error if happened
	Size(ctx context.Context, d Dir) (Result, error)
}

// sizer implement the DirSizer interface
type sizer struct {
	maxWorkersCount int // number of workers for asynchronous run
	semaphoreChan   chan struct{}
}

// NewSizer returns new DirSizer instance
func NewSizer() DirSizer {
	sizer := sizer{maxWorkersCount: 10}
	sizer.semaphoreChan = make(chan struct{}, 10)
	return &sizer
}

func (a *sizer) Size(ctx context.Context, d Dir) (Result, error) {
	result := Result{}
	resultChan := make(chan Result)
	defer close(resultChan)

	var totalError error = nil
	dirs, files, err := d.Ls(ctx)
	if err != nil {
		return result, err
	}
	var wg sync.WaitGroup
	wg.Add(len(dirs) + len(files))

	for _, file := range files {
		a.semaphoreChan <- struct{}{}
		go func(file File) {
			defer func() {
				<-a.semaphoreChan
				wg.Done()
			}()
			size, err := file.Stat(ctx)
			if err != nil {
				totalError = err
				return
			}
			atomic.AddInt64(&result.Size, size)
			atomic.AddInt64(&result.Count, 1)
		}(file)
	}

	for _, dir := range dirs {
		a.semaphoreChan <- struct{}{}
		go func(dir Dir) {
			defer func() {
				<-a.semaphoreChan
				wg.Done()
			}()
			dirResult, err := a.Size(ctx, dir)
			if err != nil {
				totalError = err
			}
			resultChan <- dirResult
		}(dir)
	}

	for i := 0; i < len(dirs); i++ {
		select {
		case <-ctx.Done():
			return result, totalError
		case dirResult := <-resultChan:
			atomic.AddInt64(&result.Size, dirResult.Size)
			atomic.AddInt64(&result.Count, dirResult.Count)
		}
	}
	wg.Wait()
	return result, totalError
}
