package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main(){
	//semaphores
	pool:=semaphore.NewWeighted(2)
	//max capacity of the instance - 2 files
	go swim("A",pool)
	go swim("B",pool)
	go swim("C",pool)
	go swim("D",pool)
	time.Sleep(5*time.Second)
	fmt.Println("main func exited")
}	

func swim(name string,pool *semaphore.Weighted){
	log.Printf("%v i want to swim",name)
	ctx:=context.Background()//context package provides way to acess propogate through all the api calls
	if err := pool.Acquire(ctx,1);err!=nil{
		fmt.Println("no lane available to swim")
	}
	fmt.Println(name)
	time.Sleep(time.Second)
	pool.Release(1)

}