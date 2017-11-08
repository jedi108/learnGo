package main

import "fmt"

type IComponent interface {
	Method() string
}

type Component struct {
}

func (component *Component) Method() string {
	return "component"
}

type Decorator struct {
	component IComponent
}

func (decorator *Decorator) Method() string {
	return "__ " + decorator.component.Method() + " __"
}

func main() {
	dec := &Decorator{&Component{}}
	fmt.Println(dec.Method())
}
