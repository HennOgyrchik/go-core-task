package main

import (
	"fmt"
	"math"
)

func main() {

	f := make(chan uint8)
	s := conveyer(f)

	go func() {
		for i := 2; i < 5; i++ {
			f <- uint8(i)
		}
		close(f)
	}()

	for v := range s {
		fmt.Println(v)
	}
}

func conveyer(in chan uint8) <-chan float64 {
	ch := make(chan float64)

	go func() {
		for a := range in {
			ch <- math.Pow(float64(a), 3)
		}
		close(ch)
	}()

	return ch
}
