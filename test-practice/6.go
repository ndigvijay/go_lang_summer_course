package main

import (
	"fmt"
	"time"
)

func printNumbers(from, to int) {
	for i := from; i <= to; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers(1, 5)
	go printNumbers(6, 10)

	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed")
}
