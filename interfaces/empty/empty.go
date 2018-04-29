package main

import (
	"fmt"
	"strconv"
)

func (myStruct *myStruct) String() string {
	return "My int value:" + strconv.Itoa(myStruct.myint)
}

func main() {
	fmt.Println("print", &myStruct{myint: 1})
	fmt.Printf("Raw: %#v\n", &myStruct{myint: 1})
	fmt.Printf("Raw: %s\n", &myStruct{myint: 1})
}
