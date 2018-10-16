package main

import "fmt"

type Costing interface {
	Cost() uint
}

type coffeBlack struct {
}

func NewCoffeBlack() *coffeBlack {
	return &coffeBlack{}
}

func (b *coffeBlack) Cost() uint {
	return 100
}

type coffeGreen struct{}

func NewCoffeGreen() *coffeGreen {
	return &coffeGreen{}
}

func (g *coffeGreen) Cost() uint {
	return 200
}

type sugar struct {
	parent *Costing
}

func NewSugar(cmp Costing) Costing {
	return sugar{
		parent: &cmp,
	}
}

func (g sugar) Cost() uint {
	return (*g.parent).Cost() + 15
}

type milk struct {
	parent *Costing
}

func NewMilk(cmp Costing) Costing {
	return milk{
		parent: &cmp,
	}
}

func (g milk) Cost() uint {
	return (*g.parent).Cost() + 15
}

func main() {

	var coffe Costing = NewCoffeBlack()
	addid := NewSugar(coffe)
	addid = NewMilk(addid)
	cost := addid.Cost()
	fmt.Println(cost)

}
