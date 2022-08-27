package concurrent

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	Go(func() {
		panic(1)
	})
	<-time.After(time.Second)
}
