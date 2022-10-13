package main

import (
	"fmt"
)

func main() {
	stopCh := make(chan struct{})
	dataCh := make(chan int)

	go func(channelCh <-chan struct{}, dataCh chan<- int) {
		val := 0
		for {
			select {
			case <-channelCh:
				close(dataCh)
				return
			case dataCh <- val:
				val++
			}
		}
	}(stopCh, dataCh)

	for curval := range dataCh {
		fmt.Println("read", curval)
		if curval > 10 {
			stopCh <- struct{}{}
			break
		}
	}
}
