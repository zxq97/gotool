package concurrent

import (
	"context"
	"sync"
)

type CtxWaitGroup struct {
	ctx  context.Context
	wg   sync.WaitGroup
	done chan struct{}
}

func NewCtxWaitGroup(ctx context.Context) *CtxWaitGroup {

}
