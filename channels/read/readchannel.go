package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	go func(out chan<- int) {
		for i := 1; i <= 10; i++ {
			// time.Sleep(time.Second * 1)
			fmt.Println("before ", i)
			out <- i
			fmt.Println("After ", i)
		}
		close(out)
	}(in)

	for i := range in {

		fmt.Println("\tget", i)
	}

	//after close channel (default):
	fmt.Println("\tget", <-in)
	fmt.Println("\tget", <-in)
	fmt.Println("\tget", <-in)
}
