package sequence

import (
	"log"
	"testing"
)

func TestChunks(t *testing.T) {
	x := make([]int, 105)
	for k := range x {
		x[k] = k
	}
	xx := Chunks[int](x, 10)
	log.Println(xx)
}
