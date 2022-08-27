package cast

import (
	"log"
	"testing"
)

func TestAtoi(t *testing.T) {
	x := Atoi("123", 0)
	log.Printf("%d %T\n", x, x)
	y := Atoi("abc", 0)
	log.Printf("%d %T\n", y, y)
}

func TestItoa(t *testing.T) {
	x := Itoa(123)
	log.Printf("%s %T\n", x, x)
}

func TestFormatInt(t *testing.T) {
	x := FormatInt(123)
	log.Printf("%s %T\n", x, x)
}

func TestParseInt(t *testing.T) {
	x := ParseInt("123", 0)
	log.Printf("%d %T\n", x, x)
	y := ParseInt("abc", 0)
	log.Printf("%d %T\n", y, y)
}
