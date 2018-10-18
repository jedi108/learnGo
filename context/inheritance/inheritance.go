package main

import (
	"context"
	"fmt"
	"time"
)

type ending interface {
	theEnd()
}

type workerChild struct {
	ctx context.Context
	num uint
}

func NewWorkerChild(ctx context.Context, num uint) *workerChild {
	newctx, _ := context.WithTimeout(ctx, time.Nanosecond*10)
	return &workerChild{ctx: newctx, num: num}
}

func (child *workerChild) theEnd() {
	<-child.ctx.Done()
	echoOfEnd(child.num, "child")
}

type WorkerCancel struct {
	workerChild *workerChild
	ctx         context.Context
	num         uint
}

func NewWorkerCancel(ctx context.Context, child *workerChild, num uint) *WorkerCancel {
	newctx, _ := context.WithTimeout(ctx, time.Hour*100)
	return &WorkerCancel{ctx: newctx, workerChild: child, num: num}
}

func (worker *WorkerCancel) theEnd() {
	<-worker.ctx.Done()
	echoOfEnd(worker.num, "worker")
}

func echoOfEnd(n uint, workName string) {
	fmt.Printf("worker name %s, num %d\n", workName, n)
}

func main() {
	ctx, finish := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		wc := NewWorkerChild(ctx, uint(i))
		wk := NewWorkerCancel(ctx, wc, uint(i))
		go wk.theEnd()
		go wc.theEnd()
	}

	time.Sleep(time.Second * 1)
	finish()
	time.Sleep(time.Second)
}
