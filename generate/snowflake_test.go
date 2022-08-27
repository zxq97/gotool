package generate

import (
	"log"
	"testing"
)

func TestIDGen_Gen(t *testing.T) {
	idg, err := NewSnowIDGen("test")
	if err != nil {
		t.Error(err)
	}
	b := 20000
	set := make(map[int64]struct{}, b)
	for i := 0; i < b; i++ {
		id := idg.Gen()
		if _, ok := set[id]; !ok {
			set[id] = struct{}{}
		} else {
			log.Fatalln("duplicate id", id)
		}
	}
}
