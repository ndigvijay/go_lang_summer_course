package main

import (
	"fmt"
	"time"
)

func execute(n string){
	for i:=0;i<10;i++{
		fmt.Println(n,i)
		time.Sleep(time.Millisecond*1000)
	}
}

func main(){
	go execute("First")
	go execute("second")
	fmt.Println("Main exited")// main doesnt wait
}