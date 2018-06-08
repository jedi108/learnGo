package main

import (
	"fmt"
)

var ii = 0

func fibo(i int) (int, int) {
	ii++
	if i == 0 || i == 1 {
		return 1, ii
	}
	fibo1, _ := fibo(i - 1)
	fibo2, _ := fibo(i - 2)
	return fibo1 + fibo2, ii
}

func main() {
	for i := 0; i < 10; i++ {
		ii = 0
		fb, iter := fibo(i)
		fmt.Printf("for = %v, fibo = %v, iter = %v\n", i, fb, iter)
	}
}
