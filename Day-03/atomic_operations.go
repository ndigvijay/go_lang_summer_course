// atomic
// data type
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	// fmt.Println("hello world")
	var ops atomic.Uint64
	var wg sync.WaitGroup
	// int ops can be accessed and modified automatically across multiple go routines
	for i := 0; i < 50 ; i++ {
		wg.Add(1)
		go func(){
			for i := 0; i < 100; i++ {
				ops.Add(1) //atomic increments the ops variable
				//single indivisible operation
			}
			wg.Done()
		}()
	}
	fmt.Println()
	wg.Wait()
}