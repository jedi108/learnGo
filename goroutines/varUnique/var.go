package main

import (
	"fmt"
	"sync"
)

type Task interface{}

type Executor interface {
	Run(t Task) error
}

func ProcessTasks(exec []Executor, tasks []Task) {
	// n executors and m>>>n tasks
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	var countExecutors int
	countExecutors = len(exec)

	i := 0
	for _, taskz := range tasks {
		go func(z int, tasks Task) {
			defer wg.Done()
			exec[z].Run(tasks)
		}(i, taskz)
		i++
		if i >= countExecutors {
			i = 0
		}
	}
	wg.Wait()
	fmt.Println("Done")
}

type MyRunner struct {
}

func (myRunner *MyRunner) Run(t Task) error {
	fmt.Println("run task", t)
	return nil
}

func main() {
	var tasks []Task
	tasks = append(tasks, 1)
	tasks = append(tasks, 2)
	tasks = append(tasks, 3)
	tasks = append(tasks, 4)
	tasks = append(tasks, 5)
	tasks = append(tasks, 6)
	MyRunner1 := &MyRunner{}
	MyRunner2 := &MyRunner{}
	var runners []Executor
	runners = append(runners, MyRunner1)
	runners = append(runners, MyRunner2)
	ProcessTasks(runners, tasks)
}
