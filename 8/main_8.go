package main

import (
	"fmt"
	"go_core_task/8/waitGroup"
	"time"
)

func main() {

	wg := waitGroup.New()

	n := 3
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println("closed", i)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ok")
}
