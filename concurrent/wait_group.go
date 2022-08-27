package concurrent

import (
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

func (wg *WaitGroup) Wait() {
	wg.wg.Wait()
}
