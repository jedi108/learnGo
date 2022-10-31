package main

import (
	"context"
	"fmt"
	"hash/fnv"
	"log"
	"time"
)

type message struct {
	Name string
}

type handler func(message2 message) error

type typeChannelCast map[int]chan message

type Broadcaster struct {
	channels typeChannelCast
	mod      int
}

func New(ctx context.Context, len, chanCap int, handler handler) Broadcaster {
	channels := make(map[int]chan message, len)
	for i := 0; i < len; i++ {
		channels[i] = make(chan message, chanCap)
		fmt.Println(i, "created")
		worker(ctx, i, channels[i], handler)
	}

	return Broadcaster{
		channels: channels,
		mod:      len,
	}
}

func worker(ctx context.Context, n int, m chan message, handler handler) {
	go func(m chan message) {
		for {
			select {
			case msg := <-m:
				if err := handler(msg); err != nil {
					log.Println(err)
				}
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					log.Println(err)
				}
				return
			}
		}
	}(m)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (b Broadcaster) Send(sku string, body interface{}) {
	chanNumber := int(hash(sku)) % b.mod
	b.channels[chanNumber] <- message{}
}

func sender(ctx context.Context) <-chan message {
	msg := make(chan message, 10)
	i := 0

	go func() {
		for {

			select {
			case <-ctx.Done():
				fmt.Println("done sender", ctx.Err())
				return
			default:
				msg <- message{Name: fmt.Sprintf("name %d", i)}
				i++
			}
		}
	}()

	return msg
}

func workers(ctx context.Context, countRun int, data <-chan message, handler func(message2 message) error) {
	for i := 0; i <= countRun; i++ {
		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("done worker", ctx.Err())
					return
				case value := <-data:
					if err := handler(value); err != nil {
						fmt.Println("err", err)
					}
				}
			}
		}(i)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Microsecond*1)

	workers(ctx, 2, sender(ctx), func(message2 message) error {
		//fmt.Println("message", message2)
		fmt.Print(".")
		return nil
	})

	time.Sleep(time.Second * 2)
	//time.Sleep(time.Second * 3)

	cancel()
	time.Sleep(time.Second * 1)
}
