/*
	Аукцион
	1. Дана макс ставка
	2. сделанно определенное количество ставок
	3. прошло определенное время
*/
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Lot struct {
	sync.Mutex
	MaxPrice   int
	MaxBids    int
	CurrentBid int
	currentCnt int
	PlayerId   int
}

func (b *Lot) GetCurrentBid() int {
	b.Lock()
	defer b.Unlock()
	return b.CurrentBid
}

func (b *Lot) SetNewBid(newBid PlayerBid) bool {
	b.Lock()
	defer b.Unlock()
	if newBid.Bid > b.CurrentBid {
		fmt.Printf("new bid: %+v\n", newBid)

		b.CurrentBid = newBid.Bid
		b.PlayerId = newBid.PlayerId
		b.currentCnt++

		// Заканчиваем аукцион если назначена максимальное количество попыток
		if b.MaxBids <= b.currentCnt {
			fmt.Println("finish by count")
			return true
		}
		// Заканчивается когда превышено максимальное количество ставок
		if b.MaxPrice <= newBid.Bid {
			fmt.Println("finish by bid")
			return true
		}
	}
	return false
}

type PlayerBid struct {
	Bid      int
	PlayerId int
}

const (
	AvgSleep   = 2000
	AvgBidStep = 10
)

func makeBid(ctx context.Context, playerId int, lot *Lot, bids chan PlayerBid) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(AvgSleep)))
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		myBid := lot.GetCurrentBid() * rand.Intn(AvgBidStep)

		select {
		case <-ctx.Done():
			return
		case bids <- PlayerBid{PlayerId: playerId, Bid: myBid}:
		default:
		}
		time.Sleep(time.Microsecond * time.Duration(rand.Intn(AvgSleep)))
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	lot := &Lot{CurrentBid: 1, MaxBids: 10, MaxPrice: 1000000}

	// Канал для ставок
	bids := make(chan PlayerBid)

	// context.Background() - Пустой контекст
	// WithTimeout - конекст по времени
	ctx, finish := context.WithTimeout(context.Background(), time.Second*5)
	// ctx, finish = context.Background()
	// ctx, finish = context.WithCancel(context.Background())

	//5 игроков которые будут делать ставки
	for i := 0; i < 5; i++ {
		go makeBid(ctx, i, lot, bids)
	}

LOOP:
	for {
		select {
		case bid := <-bids:
			if lot.SetNewBid(bid) {
				// Фуекмця отмены контекста
				finish()
				break LOOP
			}
		case <-ctx.Done():
			fmt.Println("Done:", ctx.Err())
			break LOOP
		}
	}
}
