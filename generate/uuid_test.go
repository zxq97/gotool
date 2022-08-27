package generate

import (
	"log"
	"testing"
)

func TestUUIDStr(t *testing.T) {
	log.Println(UUIDStr())
}

func TestUUIDUint(t *testing.T) {
	log.Println(UUIDUint())
}

func TestUUIDLong(t *testing.T) {
	log.Println(UUIDLong())
}
