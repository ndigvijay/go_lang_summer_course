package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup){
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	fmt.Println("counter",counter)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("main exited")

}