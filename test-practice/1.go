package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 done")
	}()

	//wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2 done")
	}()

	wg.Wait()
	fmt.Println("All goroutines finished")
}
