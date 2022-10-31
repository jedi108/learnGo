package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Item struct {
	name string
}

func Fetch(ctx context.Context, name string) (Item, error) {

}

func consumse(a, b Item) {

}

func main() {
	var a, b Item
	var ctx = context.TODO()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		a, err := Fetch(ctx, "a")
		return err
	})
	g.Go(func() error {
		b, err := Fetch(ctx, "a")
		return err
	})

	consume(a, b)
}
