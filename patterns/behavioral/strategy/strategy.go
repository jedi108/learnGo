package main

import "fmt"

type StrategyAlgs interface {
	Order()
}

type DataAlgoritmBuy struct {
}

func (dataAlg *DataAlgoritmBuy) Order() {
	fmt.Println("buy")
}

type DataAlgoritmSell struct {
}

func (dataAlg *DataAlgoritmSell) Order() {
	fmt.Println("sell")
}

type Context struct {
	algoritmData StrategyAlgs
}

func (context *Context) SetAlgoritm(alg StrategyAlgs) {
	context.algoritmData = alg
}

func (context *Context) ExecuteAlgoritm() {
	context.algoritmData.Order()
}

func main() {
	context := &Context{}
	context.SetAlgoritm(&DataAlgoritmBuy{})
	context.ExecuteAlgoritm()
	context.SetAlgoritm(&DataAlgoritmSell{})
	context.ExecuteAlgoritm()
}
