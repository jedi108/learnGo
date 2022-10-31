package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestBroadcaster_Send(t *testing.T) {
	from := 32768
	for i := 32768; i <= from*100; i++ {
		result := i % 100
		if result > 100 {
			t.Fatal("sex")
		}
		fmt.Println(result)
	}
}

func TestBlockReader(t *testing.T) {
	count := 5
	var i uint32
	b := New(context.TODO(), count, 1, func(message2 message) error {
		atomic.AddUint32(&i, 1)
		if atomic.LoadUint32(&i) > 1 {
			//time.Sleep(time.Millisecond * time.Duration(math.Pow(float64(i), float64(i))) * time.Duration(1000))
			atomic.StoreUint32(&i, 0)
		}
		fmt.Println("msg <- ", message2)
		return nil
	})

	for i := 32867; i < 32867+100; i++ {
		hash := i % count
		cham, ok := b.channels[hash]
		if ok {
			fmt.Println("-> send:", i, hash)
			cham <- message{Name: fmt.Sprintf("A %d; hash = %d", i, hash)}
		} else {
			t.Fatalf("error send to: %v, %d, %d \n", hash, i, count)
		}
	}
}

func TestGO(t *testing.T) {
	c := make(chan int, 1)
	for i := 0; i < 2; i++ {
		go func(n int) {
			for {
				select {
				case cc := <-c:
					fmt.Println("CC", n, cc)
					if cc > 2 {
						time.Sleep(time.Millisecond * time.Duration((cc)))
					}
				}
			}
		}(i)
	}

	for i := 0; i < 100; i++ {
		c <- i
	}
}
