// https://4gophers.ru/articles/modeli-konkurentnosti-v-go/

package main

import (
	"fmt"
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

func main() {
	// Создаем необходимые каналы и выводим значения.
	for n := range sq(sq(gen(1, 2, 3))) {
		fmt.Println(n) // 16 затем 81
	}
}

// func main() {
// 	// Создаем необходимые каналы.
// 	c := gen(1, 2, 3)
// 	out := sq(c)

// 	// Выводим значения.
// 	fmt.Println(<-out)
// 	fmt.Println(<-out)
// 	fmt.Println(<-out)
// }
