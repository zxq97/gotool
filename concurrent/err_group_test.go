package concurrent

import (
	"context"
	"log"
	"testing"
)

func TestErrGroup_Go(t *testing.T) {
	eg := NewErrGroup(context.TODO())
	eg.Go(func() error {
		panic(1)
	})
	eg.Go(func() error {
		log.Println(1)
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}
	log.Println("done")
}
