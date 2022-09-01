package concurrent

import (
	"context"
	"log"
	"runtime/debug"
	"sync"
)

type CtxWaitGroup struct {
	ctx  context.Context
	wg   sync.WaitGroup
	err  error
	done chan struct{}
}

func NewCtxWaitGroup(ctx context.Context) *CtxWaitGroup {
	return &CtxWaitGroup{
		ctx:  ctx,
		wg:   sync.WaitGroup{},
		done: make(chan struct{}),
	}
}

func (cwg *CtxWaitGroup) Go(fn func()) {
	cwg.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, string(debug.Stack()))
			}
			cwg.wg.Done()
		}()
		fn()
	}()
}

func (cwg *CtxWaitGroup) Wait() error {
	go func() {
		cwg.wg.Wait()
		close(cwg.done)
	}()
	select {
	case <-cwg.done:
		return nil
	case <-cwg.ctx.Done():
		return cwg.ctx.Err()
	}
}
