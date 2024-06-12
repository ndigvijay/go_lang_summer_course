package main

import (
	"fmt"
	"sync"
)

var counter int
func sayHello(wg *sync.WaitGroup,n string){
	counter++ // shared variable , not protected
	fmt.Println("hello ",n,"counter",counter)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go sayHello(&wg,"alice")
	go sayHello(&wg,"bob")
	wg.Wait()
	fmt.Println("main is done with its job")
}