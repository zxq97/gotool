package concurrent

import (
	"log"
	"runtime/debug"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, string(debug.Stack()))
			}
		}()
		fn()
	}()
}
