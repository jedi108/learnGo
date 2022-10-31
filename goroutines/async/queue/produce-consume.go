package main

import "fmt"

type Item struct {
	name string
}

// Glob finds item with names matching patter and send them on the returner channel& It closed the channel when all items have been sent
func Glob(pattern string) <-chan Item {
	// A channel fed by one gouroutine and read by another acts as a queue
	c := make(chan Item)
	go func() {
		defer close(c)
		for {
			c <- Item{name: pattern}
		}
	}()
	return c
}

func main() {
	for item := range Glob("[ab]*") {
		fmt.Println(item)
	}
}
