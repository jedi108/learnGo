package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("init instance")
		instance = &singleton{}
	})
	return instance
}

func main() {
	GetInstance()
	GetInstance()
}
