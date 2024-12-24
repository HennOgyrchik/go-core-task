package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := make([]chan int, 0, 10)

	for i := 0; i < cap(arr); i++ {

		ch := make(chan int)
		arr = append(arr, ch)

		go func(c chan int) {
			for j := 0; j < 3; j++ {
				c <- i
			}
			close(ch)
		}(ch)

	}

	m := merge(arr)

	for v := range m {
		fmt.Printf("%v ", v)
	}

}

func merge[T comparable](arr []chan T) <-chan T {

	var wg sync.WaitGroup
	wg.Add(len(arr))

	mainCh := make(chan T)

	for _, ch := range arr {
		go func() {

			for v := range ch {
				mainCh <- v
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(mainCh)
	}()

	return mainCh
}
