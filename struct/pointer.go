package main

import "fmt"

type structA struct {
	state int
}

func (s structA) setA(arg int) {
	s.state = arg
}

type structB struct {
	state int
}

func (s *structB) setA(arg int) {
	s.state = arg
}

func main() {
	a := &structA{}
	b := structB{}

	a.setA(1)
	b.setA(2)

	fmt.Println(a.state)
	fmt.Println(b.state)
}
