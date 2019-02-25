//Mutual Exclusion
// 		A concurrent process holds exclusive rights to a resource at any one time.
// Wait For Condition
// 		A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
// No Preemption
// 		A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
// Circular Wait
// 		A concurrent process (P1) must be waiting on a chain of other concurrent pro‐ cesses (P2), which are in turn waiting on it (P1), so it fulfills this final condition too.

package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func (v value) lock() {
	v.mu.Lock()
}

func (v value) unlock() {
	v.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	printSub := func(i int, v1, v2 *value) {
		defer wg.Done()
		fmt.Println("lock A ", i)
		v1.lock()
		v1.value++
		defer func() {
			fmt.Println("unlock A", i)
			v1.unlock()
		}()

		time.Sleep(time.Microsecond * 2)

		fmt.Println("lock B ", i)

		v2.lock()
		v2.value++
		defer func() {
			fmt.Println("unlock B ", i)
			v2.unlock()
		}()

		fmt.Printf("sum = %v; g=%d\n", v1.value+v2.value, i)
	}

	var a, b value

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			printSub(i, &a, &b)
		}(i)
	}
	wg.Wait()
}
