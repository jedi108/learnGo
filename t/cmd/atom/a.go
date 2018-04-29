package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var a int32
	var wg sync.WaitGroup
	for i := 0; i < 150; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&a, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
