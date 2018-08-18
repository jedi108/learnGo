package main

import (
	"context"
	"log"
	"time"
)

func main() {
	defer log.Println("end main")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer log.Println("cancel")
		time.Sleep(time.Second * 1)
		cancel()
	}()
	go op1(ctx)
	// <-ctx.Done()
}

func op1(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("done in op1")
	}
}
