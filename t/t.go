package main

import (
	"strconv"
	"fmt"
)

type Obj struct{
	Value string
}

func main() {
	m := make(map[string]Obj)
	for i:=0;i<10;i++ {
		a := &Obj{}
		a.Value = strconv.Itoa(i)
		m[a.Value] = a
	}
	for _, v := range m {
		(v.Value) = "aaa"
		v.Value = "new_value"
		fmt.Print("s")
	}
	for _, v := range m {
		fmt.Println(v.Value)
	}
}