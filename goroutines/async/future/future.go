package main

import "fmt"

type Item struct {
	name string
}

//Future pattern
func fetch(name string) <-chan Item {
	c := make(chan Item, 1)
	go func() {
		/// [....]
		item := Item{name: name}
		c <- item
	}()
	return c
}

func consume(one Item, two Item) {
	fmt.Printf("one: %#v; two: %#v", one, two)
}

func main() {
	a := fetch("a")
	b := fetch("b")
	consume(<-a, <-b)
}
