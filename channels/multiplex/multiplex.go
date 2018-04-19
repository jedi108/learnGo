package main

import (
	"fmt"
)

func main() {
	channelCh := make(chan struct{})
	dataCh := make(chan int)

	go func(channelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-channelCh:
				return
			case dataCh <- val:
				val++
			}
		}
	}(channelCh, dataCh)

	for curval := range dataCh {
		fmt.Println("read", curval)
		if curval > 3 {
			channelCh <- struct{}{}
			break
		}
	}
}
