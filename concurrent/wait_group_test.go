package concurrent

import (
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
