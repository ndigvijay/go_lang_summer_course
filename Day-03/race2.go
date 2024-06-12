package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup,n string){
	fmt.Println("hello "+n)
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