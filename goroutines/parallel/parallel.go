package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 128; i++ {
			fmt.Print("- ")
		}
	}()

	go func() {
		defer wg.Done()
		for z := 1; z < 128; z++ {
			fmt.Print("+ ")
		}
	}()
	wg.Wait()
	fmt.Println("", time.Since(start))
}
