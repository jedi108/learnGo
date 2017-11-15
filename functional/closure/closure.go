package main

import (
	"fmt"
)

func closureCounter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	addCount := closureCounter()
	fmt.Println(addCount())
	fmt.Println(addCount())
	fmt.Println(addCount())
}
