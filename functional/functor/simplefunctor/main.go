package main

import "fmt"

type intSliceFunctor struct {
	intSlice   []int
	mapperFunc func(int) int
}

func liftIntSlice(someSlice []int) *intSliceFunctor {
	return &intSliceFunctor{intSlice: someSlice}
}

func (functor *intSliceFunctor) maps(f func(int) int) *intSliceFunctor {
	for i, v := range functor.intSlice {
		functor.intSlice[i] = f(v)
	}
	return functor
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	functor := liftIntSlice(intSlice)

	mapperFunc := func(i int) int {
		return i + 10
	}

	mapped := functor.maps(mapperFunc)
	fmt.Println(mapped.intSlice)
}
