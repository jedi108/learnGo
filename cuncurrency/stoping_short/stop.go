package main

import (
	"fmt"
	"sync"
)

func gen1(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3, 4, 5, 6, 7)

	c1 := sq(in)
	c2 := sq(in)

	// Получаем первое значение из выходного канала.
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 or 9
	return
	// Так как мы не получаем второе значение и `out`,
	// то рутина зависает при попытке отправки чего либо.
}
