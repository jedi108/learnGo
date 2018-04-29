package main

import "fmt"
import "sync"

import "time"

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer func() {
		wg.Done()

	}()
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go worker(w, jobs, results)
	}

	for j := 1; j <= 90; j++ {
		jobs <- j
	}

	for msg := range results {
		fmt.Println("worker", msg)
	}

	wg.Wait()
	close(jobs)
	close(results)
	fmt.Println("main finished")

}
