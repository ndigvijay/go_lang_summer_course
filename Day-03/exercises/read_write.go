package main

import (
	"fmt"
	"sync"
)


var shared int = 70 
var mutex sync.Mutex
var wg sync.WaitGroup
func main(){
	wg.Add(2)
	//write
	go func(){
		mutex.Lock()
		shared++;
		wg.Done()
		mutex.Unlock()
	}()
	// read
	go func(){
		mutex.Lock()
		fmt.Println("reading shared",shared)
		mutex.Unlock()
		wg.Done()

	}()
	wg.Wait()
	fmt.Println(shared)
}