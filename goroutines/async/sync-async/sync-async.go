package main

import "fmt"

type In string
type Out string

func Async(x In) <-chan Out {
	c := make(chan Out, 1)
	go func() {
		c <- Sync(x)
	}()
	return c
}

func Sync(x In) Out {
	fmt.Println(x)
	c := Async(x)
	return <-c
}

func main() {
	a := Sync("A")
	fmt.Println(a)
}
