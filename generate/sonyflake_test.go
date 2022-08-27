package generate

import (
	"log"
	"testing"
)

func TestSonyIDGen_GenID(t *testing.T) {
	idg, err := NewSonyIDDen("2022-01-01 00:00:00")
	if err != nil {
		t.Error(err)
	}
	b := 20000
	set := make(map[int64]struct{}, b)
	for i := 0; i < b; i++ {
		id := idg.GenID()
		if _, ok := set[id]; !ok {
			set[id] = struct{}{}
		} else {
			log.Fatalln("duplicate id", id)
		}
	}
}
