package main

import (
	"fmt"
	"time"
)

func sendToChannel(c chan<- string, msg string, delay time.Duration) {
	time.Sleep(delay)
	c <- msg
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go sendToChannel(channel1, "message from channel 1", 2*time.Second)
	go sendToChannel(channel2, "message from channel 2", 1*time.Second)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received:", msg2)
		}
	}

	fmt.Println("All messages received")
}
