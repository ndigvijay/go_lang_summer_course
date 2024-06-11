// go routine

// subroutine in c

// achieve concurrency

// construct

// execute mulitple loops at the same time

// looping parellel

package main

import (
	"fmt"
	"time"
)

func ndv(){
	for i:=0;i<5;i++{
		time.Sleep(2*time.Second)
		fmt.Println("yo")
		// time.Sleep(5*time.Second)
	}
}

func main(){
	go ndv() // lightweight thread 
	for i:=0;i<5;i++{
		time.Sleep(2*time.Second)
		fmt.Println("yo")
	}
}

//multiplexing

// write a code to show go routine in go(any example)