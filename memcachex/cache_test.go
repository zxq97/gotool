package memcachex

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestMemcacheX_SetCtx(t *testing.T) {
	mcx := NewMemcacheX([]string{"127.0.0.1:11211"})
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	err := mcx.SetCtx(ctx, "k", []byte("1234"), 3600)
	if err != nil {
		log.Println(err)
	}
}

func TestMemcacheX_GetCtx(t *testing.T) {
	mcx := NewMemcacheX([]string{"127.0.0.1:11211"})
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	val, err := mcx.GetCtx(ctx, "k")
	log.Println(val, err)
}

func TestMemcacheX_GetMultiCtx(t *testing.T) {
	mcx := NewMemcacheX([]string{"127.0.0.1:11211"})
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	val, err := mcx.GetMultiCtx(ctx, []string{"k"})
	log.Println(val, err)
}
