// Вопрос со звездочкой: Какая из строк (A или B) позволит получить лучшее время выполнения при условии что у нас 8 ядер (numcpu=8)?
// https://habr.com/ru/post/195574/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("NumCPU", numCPU)

	runtime.GOMAXPROCS(numCPU) // A
	//runtime.GOMAXPROCS(1) // B

	ch1 := make(chan int)
	ch2 := make(chan float64)

	go func() {
		for i := 0; i < 1000000; i++ {
			ch1 <- i
		}
		ch1 <- -1
		ch2 <- 0.0
	}()
	go func() {
		total := 0.0
		for {
			t1 := time.Now().UnixNano()
			for i := 0; i < 100000; i++ {
				m := <-ch1
				if m == -1 {
					ch2 <- total
				}
			}
			t2 := time.Now().UnixNano()
			dt := float64(t2-t1) / 1000000.0
			total += dt
			fmt.Println(dt)
		}
	}()

	fmt.Println("Total:", <-ch2, <-ch2)
}
