package main

import (
	"fmt"
	"time"
)

func main(){
	ch1:=make(chan string,10)
	ch2:=make(chan string,20)
	go func(){
		time.Sleep(1*time.Second)
		ch1<-"hi"
	}()
	go func(){
		time.Sleep(2*time.Second)
		ch2<-"hello"
	}()
	for i:=0;i<2;i++{
		select{
		case msg1:=<-ch1:
		fmt.Println(msg1)
		case msg2:=<-ch2:
		fmt.Println(msg2)
		}
	}

}