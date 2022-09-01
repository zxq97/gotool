package concurrent

import (
	"context"
	"log"
	"runtime/debug"
	"sync"
)

type WaitGroup struct {
	wg sync.WaitGroup
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{wg: sync.WaitGroup{}}
}

func (wg *WaitGroup) Go(fn func()) {
	wg.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, string(debug.Stack()))
			}
			wg.wg.Done()
		}()
		fn()
	}()
}

func (wg *WaitGroup) GoC(ctx context.Context, fn func()) {
	done := make(chan struct{})
	wg.wg.Add(1)
	go func() {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err, string(debug.Stack()))
				}
			}()
			fn()
		}()
		select {
		case <-done:
		case <-ctx.Done():
		}
		wg.wg.Done()
	}()
}

func (wg *WaitGroup) Wait() {
	wg.wg.Wait()
}
