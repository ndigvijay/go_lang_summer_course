package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
}
