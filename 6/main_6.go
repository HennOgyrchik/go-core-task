package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	randCh := randGen(ctx, 1, 10)

	for i := 0; i < 3; i++ {
		fmt.Println(<-randCh)
	}

	cancel()

}

func randGen(ctx context.Context, from, to int) chan int {
	ch := make(chan int)

	go func() {
		rand.New(rand.NewSource(time.Now().UnixNano()))

		for {
			select {
			case ch <- rand.Intn(to-from) + from:
			case <-ctx.Done():
				close(ch)
				return
			}
		}
	}()

	return ch

}
