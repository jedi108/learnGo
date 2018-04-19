package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
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

	// Запуск go-рутины для каждого входного канала из `cs`. `output`
	// копирует значения из входного канала `с` пока `с` не будет
	// закрыт. Затем вызывается `wg.Done`.
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

	// Запуск go-рутины, которая закроет `out` канал после
	// завершения всех  `output` go-рутин. Этот код должен
	// выполняться только после вызова `wg.Add`.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// Распределяем работу между двумя воркерами для считывания данных из `in`.
	c1 := sq(in)
	c2 := sq(in)

	// Объединяем вывод из c1 и c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 затем 9, или 9 затем 4
	}
}
