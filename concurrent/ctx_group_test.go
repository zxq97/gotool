package concurrent

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestCtxWaitGroup_Go(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Microsecond)
	defer cancel()
	cwg := NewCtxWaitGroup(ctx)
	cwg.Go(func() {
		<-time.After(time.Second)
		log.Println(1)
	})
	cwg.Go(func() {
		log.Println(2)
	})
	err := cwg.Wait()
	log.Println(err)
}