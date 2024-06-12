package main

import (
	"fmt"
	"log"
	"sync"
)

var mutex sync.Mutex
func counter(count *int,wg *sync.WaitGroup){
	for i:=0;i<1000;i++{
		mutex.Lock()
		*count++
		mutex.Unlock()
		log.Println(*count)
		}
	wg.Done()

}

func main(){
	var count1 int=0
	var count2 int=0
	var wg sync.WaitGroup
	wg.Add(10)
	//counter 1
	go counter(&count1,&wg)
	go counter(&count1,&wg)
	go counter(&count1,&wg)
	go counter(&count1,&wg)
	go counter(&count1,&wg)
	//counter 2
	go counter(&count2,&wg)
	go counter(&count2,&wg)
	go counter(&count2,&wg)
	go counter(&count2,&wg)
	go counter(&count2,&wg)
	wg.Wait() 
	fmt.Println(count1,count2)

}