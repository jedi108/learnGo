package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d) // simulate a workload
			cond.L.Lock() // changes must be made when cond.L is locked
			values[i] = string('a' + i)
			cond.Broadcast() // called when cond.L lock is acquired
			cond.L.Unlock()
			// The above "cond.Signal()" line can also be put here.
			//cond.Signal() // called when cond.L lock is released.
		}(i)
	}

	// This function must be called when cond.L is locked.
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}
	for !checkCondition() {
		cond.Wait() // must be called when cond.L is locked
	}
	cond.L.Unlock()
}
