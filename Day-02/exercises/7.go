package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
	SRN  string
	CGPA float64
}

func send_message(nameCh chan string, srnCh chan string, cgpaCh chan float64) {
	go func() {
		time.Sleep(5 * time.Millisecond)
		nameCh <- "John Doe"
	}()
	go func() {
		time.Sleep(5 * time.Millisecond)
		srnCh <- "SRN12345"
	}()
	go func() {
		time.Sleep(5 * time.Millisecond)
		cgpaCh <- 9.5
	}()
	for i:=0;i<3;i++{
		select{
		case msg1:=<-nameCh:
		fmt.Println(msg1)
		case msg2:=<-srnCh:
		fmt.Println(msg2)
		case msg3:=<-cgpaCh:
		fmt.Println(msg3)
		}
	}

}

func main() {
	nameCh := make(chan string)
	srnCh := make(chan string)
	cgpaCh := make(chan float64)
	send_message(nameCh, srnCh, cgpaCh)
}