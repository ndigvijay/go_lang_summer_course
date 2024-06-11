package main

import (
	"fmt"
	"time"
)

// Communication medium. Data can be transferred from source to target destination. Using channel synchronization can be acheived.
// 3 types of channels namely:
// 1. Unbuffered channel - it is a type of channel with capacity zero and stores only one value
// 2. Buffered channel - It is a type of channel with specific capacity and stores smore than one value.
// 3. Pipeline - o/p of one stage is passed as input to the next stage
// 4.

func main() {
	// Buffered channels
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 15)

	// Sending data to buffered channels
	ch1 <- "Hello"
	ch2 <- "World"
	
	// Print the first message
	fmt.Println(<-ch1)

	// Start a goroutine to count seconds
	done := make(chan bool)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Elapsed time: %d seconds\n", i)
			time.Sleep(time.Second)
		}
		done <- true
	}()

	// Main goroutine sleeps for 5 seconds
	time.Sleep(time.Second * 5)
	<-done

	// Print the second message
	fmt.Println(<-ch2)
}