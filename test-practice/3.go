package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter--
		}
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
