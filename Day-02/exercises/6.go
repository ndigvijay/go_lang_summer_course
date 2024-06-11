package main

import (
	"fmt"
	"time"
)



func send_message(ch1 chan string, ch2 chan int, ch3 chan float64) {
	go func() {
		time.Sleep(5 * time.Millisecond)
		ch1 <- "hi"
	}()
	go func() {
		time.Sleep(5 * time.Millisecond)
		ch2 <- 55
	}()
	go func() {
		time.Sleep(5 * time.Millisecond)
		ch3 <- 270.58
	}()
	for i:=0;i<3;i++{
		select{
		case msg1:=<-ch1:
		fmt.Println(msg1)
		case msg2:=<-ch2:
		fmt.Println(msg2)
		case msg3:=<-ch3:
		fmt.Println(msg3)
		}
	}
}

func main(){
	ch1:=make(chan string,10)
	ch2:=make(chan int,20)
	ch3:=make(chan float64,30)
	send_message(ch1,ch2,ch3)
	
}