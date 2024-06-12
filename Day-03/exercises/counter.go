package main

import (
	"fmt"
	"log"
	"sync"
)

func counter(count *int,wg *sync.WaitGroup){
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	for i:=0;i<10;i++{
		*count++
		log.Println(*count)
	}
	wg.Done()

}

func main(){
	var count int=0
	var wg sync.WaitGroup
	wg.Add(1)
	go counter(&count,&wg)
	wg.Wait() 
	fmt.Println(count)


}