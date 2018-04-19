package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(waorkerNum int, in <-chan string) {
	for input := range in {
		fmt.Println("worker", waorkerNum, input)
		runtime.Gosched()
	}
	fmt.Println("end wprker", waorkerNum)
}

func main() {
	workinput := make(chan string, 2)
	for i := 1; i < 3; i++ {
		go worker(i, workinput)
	}

	month := []string{"янв", "фев", "март", "апр", "май", "июнь", "июль", "авг", "сент", "окт", "ноя", "дек"}
	for _, monthName := range month {
		workinput <- monthName
	}
	// close(workinput)

	time.Sleep(time.Second)
}
