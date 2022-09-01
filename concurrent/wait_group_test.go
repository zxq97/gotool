package concurrent

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestWaitGroup_Go(t *testing.T) {
	wg := NewWaitGroup()
	wg.Go(func() {
		panic(1)
	})
	wg.Go(func() {
		<-time.After(time.Second)
	})
	wg.Wait()
}

func TestWaitGroup_GoC(t *testing.T) {
	wg := NewWaitGroup()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond)
	defer cancel()
	now := time.Now()
	wg.GoC(ctx, func() {
		<-time.After(time.Second)
	})
	wg.GoC(ctx, func() {
		log.Println(1)
	})
	wg.Wait()
	log.Println(time.Since(now))
}
