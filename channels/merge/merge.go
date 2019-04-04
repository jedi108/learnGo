package main

import (
	"fmt"
	"sync"
)

func merge3(c1 chan int, c2 chan int) (out chan int) {
	out = make(chan int)
	var (
		newFanIn = func() (instance func(in <-chan int), mergedChannel chan int) {
			wg := new(sync.WaitGroup)
			mergedChannel = make(chan int)
			go func() {
				wg.Wait()
				close(mergedChannel)
			}()

			return func(in <-chan int) {
				wg.Add(1)
				go func() {
					for c := range in {
						mergedChannel <- c
					}
					wg.Done()
				}()
			}, mergedChannel
		}

		sortAndSend = func(inputValue int, output chan<- int, buff *int) {
			switch {
			case *buff == 0 || inputValue == *buff:
				*buff = inputValue
			case inputValue < *buff:
				output <- inputValue
			case inputValue > *buff:
				output <- *buff
				*buff = inputValue
			}
		}

		sortReadWrite = func(in <-chan int, outer chan<- int) {
			temp := 0
			for m := range in {
				sortAndSend(m, outer, &temp)
			}
			outer <- temp
			close(outer)
		}

		merger, mergedChannel = newFanIn()
	)

	merger(c1)
	merger(c2)
	//merger(c3)
	//merger(c4) -- for mor in channels

	//-- simple sort
	go sortReadWrite(mergedChannel, out)
	//--

	//-- deep sort
	//sorted1Channel := make(chan int)
	//go sortReadWrite(mergedChannel, sorted1Channel)
	//go sortReadWrite(sorted1Channel, out) //--  for deep sort
	//--

	return out
}

func merge1(c1 chan int, c2 chan int) (out chan int) {
	out = make(chan int)

	var wg sync.WaitGroup
	merger := func(ch chan int, merged chan int) {
		wg.Add(1)
		go func() {
			for c := range ch {
				merged <- c
			}
			wg.Done()
		}()
	}

	merged := make(chan int)
	merger(c1, merged)
	merger(c2, merged)

	go func() {
		wg.Wait()
		close(merged)
	}()

	sortAndSend := func(inputValue int, output chan<- int, buff *int) {
		switch {
		case *buff == 0 || inputValue == *buff:
			*buff = inputValue
		case inputValue < *buff:
			output <- inputValue
		case inputValue > *buff:
			output <- *buff
			*buff = inputValue
		}
	}

	sortMerged := func(in <-chan int, outer chan<- int) {
		go func() {
			temp := 0
			for m := range in {
				sortAndSend(m, outer, &temp)
			}
			outer <- temp
			close(outer)
		}()
	}

	sortMerged(merged, out)

	return out
}

func merge2(c1 chan int, c2 chan int) (out chan int) {
	out = make(chan int)
	var ok1, ok2 = true, true

	prev := 0
	check := func(new int) {
		if prev == 0 {
			prev = new
			return
		}

		if new < prev {
			out <- new
		} else {
			out <- prev
			prev = new
		}

		if ok1 == false && ok2 == false {
			fmt.Println("close out")
			close(out)
		}
	}

	go func() {
	loops:
		for {
			select {
			case a, ok1 := <-c1:
				if ok1 == false {
					if ok2 == false {
						check(a)
						break loops
					}
					continue
				}
				check(a)
			case b, ok2 := <-c2:
				if ok2 == false {
					if ok1 == false {
						check(b)
						break loops
					}
					continue
				}
				check(b)
			}
		}
	}()

	return out
}

func send(s []int) chan int {
	out := make(chan int)

	go func() {
		for _, i := range s {
			out <- i
		}
		close(out)
	}()

	return out
}

func main() {
	s1 := []int{1, 2, 3, 10, 11}
	s2 := []int{2, 5, 6, 7, 8}

	c1 := send(s1)
	c2 := send(s2)

	out := merge3(c1, c2)

	for n := range out {
		fmt.Println("received: ", n)
	}
}
