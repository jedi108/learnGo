package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	My struct {
		wg sync.WaitGroup
		mu *sync.Mutex
		i  int64
		s  string
	}
)

func New() *My {
	return &My{}
}

func Get(my *My) {
	// my.mu.Lock()
	fmt.Print("  ", my.s, my.i)
	// my.mu.Unlock()
}

func inc(my *My) {
	//
	my.i++
	my.s = "++"
	//
}

func dec(my *My) {
	//
	my.i--
	my.s = "--"
	//
}

func (my *My) Get() {
	Get(my)
	// my.wg.Done()
}

func (my *My) inc() {
	my.mu.Lock()
	defer my.mu.Unlock()
	inc(my)
	// my.mu.Unlock()
	// my.wg.Done()
}

func (my *My) dec() {
	my.mu.Lock()
	defer my.mu.Unlock()
	dec(my)
	// my.mu.Unlock()
	// my.wg.Done()
}

func main() {
	const iterate int = 90000

	MyVal := New()
	time.Sleep(time.Second * 1)
	// MyVal.wg.Add(iterate)

	for i := 1; i <= iterate/3; i++ {
		go MyVal.inc()
		go MyVal.dec()
		go MyVal.Get()
	}

	time.Sleep(time.Second * 1)

	// MyVal.wg.Wait()
	fmt.Println("\n..........")
}
