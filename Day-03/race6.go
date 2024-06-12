package main

import (
	"fmt"
	"sync"
)
var count int=0
func worker(wg *sync.WaitGroup,mutex *sync.Mutex){
	mutex.Lock()
	count+=1
	mutex.Unlock()
	wg.Done()
	fmt.Println(count)
}

func main(){
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg,&mutex)
		wg.Wait()
	}
}