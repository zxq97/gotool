package concurrent

import (
	"context"
	"log"
	"runtime/debug"

	"golang.org/x/sync/errgroup"
)

type ErrGroup struct {
	eg  *errgroup.Group
}

func NewErrGroup(ctx context.Context) *ErrGroup {
	eg, _ := errgroup.WithContext(ctx)
	return &ErrGroup{eg: eg}
}

func (eg *ErrGroup) Wait() error {
	return eg.eg.Wait()
}

func (eg *ErrGroup) Go(fn func() error) {
	eg.eg.Go(func() error {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, string(debug.Stack()))
			}
		}()
		return fn()
	})
}
