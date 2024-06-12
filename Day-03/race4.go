package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex
func sayHello(wg *sync.WaitGroup,n string){
	mutex.Lock()
	defer mutex.Unlock() // releases the lock after accessing and modifying the counter variable
	counter++ // shared variable , not protected
	fmt.Println("hello ",n,"counter",counter)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go sayHello(&wg,"alice") // running on a different thread
	go sayHello(&wg,"bob")  // running on a different thread
	wg.Wait()
	fmt.Println("main is done with its job")
}