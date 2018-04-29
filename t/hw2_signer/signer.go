package main

import (
	"fmt"
	"runtime"
	"strconv"
)

// сюда писать код

func ExecutePipeline(jobs ...job) {
	var out []chan interface{}
	out = make([]chan interface{}, len(jobs), len(jobs))
	done := make(chan struct{})
	for i := 0; i < len(jobs); i++ {
		out[i] = make(chan interface{})
		if i%2 == 0 {
			if i == 0 {
				go func(i int) {
					defer close(out[i])
					fmt.Println("job", i)
					jobs[i](make(chan interface{}), out[i]) //[0] = in 0, out 0
					runtime.Gosched()
				}(i)
			} else {
				go func(i int) {
					fmt.Println("job", i)
					jobs[i](out[i-1], out[i]) //[2] = out 1, out 2
				}(i)
			}
		} else {
			go func(i int) {
				fmt.Println("job", i)
				jobs[i](out[i-1], out[i]) //[1] = out 0, out 1
				fmt.Println("done")
				done <- struct{}{}
			}(i)
		}
	}
	<-done
}

//SingleHash считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~),
//где data - то что пришло на вход (по сути - числа из первой функции)
// * crc32 считается через функцию DataSignerCrc32
// * md5 считается через DataSignerMd5
func SingleHash(in, out chan interface{}) {
	interf := <-in
	interfStr := strconv.Itoa(interf.(int))
	out <- DataSignerCrc32(interfStr) + "~" + DataSignerCrc32(DataSignerMd5(interfStr))
}
func MultiHash(in, out chan interface{})      {}
func CombineResults(in, out chan interface{}) {}

func main() {
	freeFlowJobs := []job{
		job(func(in, out chan interface{}) {
			out <- 1
			out <- 2
		}),
		job(func(in, out chan interface{}) {
			for i := range in {
				out <- i
			}
		}),
		job(func(in, out chan interface{}) {

			for i := range in {
				fmt.Println(i)
			}
			fmt.Println("the ")
		}),
	}
	ExecutePipeline(freeFlowJobs...)
}
