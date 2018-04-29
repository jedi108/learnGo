package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationNum  = 6
	goroutinetNum = 5
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{}
	defer wg.Done()
	for j := 0; j < iterationNum; j++ {
		// if j%2 == 0 {
		// 	<-quotaCh
		// 	quotaCh <- struct{}{}
		// }
		fmt.Print(formatWork(in, j))
		runtime.Gosched()
	}
	<-quotaCh
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in)+"#", "\t", in, j)
}

func main() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < goroutinetNum; i++ {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}
