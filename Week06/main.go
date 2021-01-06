package main

import (
	"loong.me/gopher/rolling"
)

func main() {
	n := rolling.NewNumber(10)
	for i := 0; i < 100; i++ {
		n.Increment(1)
	}
}
