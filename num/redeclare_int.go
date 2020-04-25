package main

import "fmt"

func main() {
	a := 1

	if a > 0 {
		a := 2
		a = a + 100
		a += 1
	}

	fmt.Print(a)
}
