package main

import (
	"fmt"
)

type Stater interface {
	ApplyState() string
}

type Computer struct {
	State Stater
}

func NewComputerOn() *Computer {
	return &Computer{State: &PowerOn{}}
}

func (computer *Computer) SetState(state Stater) {
	computer.State = state
}

func (computer *Computer) ApplyGetState() string {
	return computer.State.ApplyState()
}

type PowerOn struct{}
type PowerOff struct{}
type Hibernate struct{}

func (pcState *PowerOn) ApplyState() string {
	return "power on"
}
func (pcState *PowerOff) ApplyState() string {
	return "power off"
}
func (pcState *Hibernate) ApplyState() string {
	return "hibernate"
}

func main() {
	computer := NewComputerOn()
	fmt.Println(computer.ApplyGetState())
	computer.SetState(&Hibernate{})
	fmt.Println(computer.ApplyGetState())
	computer.SetState(&PowerOff{})
	fmt.Println(computer.ApplyGetState())
}
