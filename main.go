package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	for i := range gen(ctx) {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("Got from gen: %d \n", i)
		if i == 10 {
			cancel()
		}
	}
	time.Sleep(2 * time.Second)
}

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	var i int

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("closing goroutines...")
				fmt.Println(ctx.Err())
				close(out)
				return
			case out <- i:
				i++
			}
		}
	}()

	return out
}
