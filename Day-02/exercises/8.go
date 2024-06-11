package main

import (
	"fmt"
)

func main(){
	// fmt.Println("hello world")
	var name string="digvijay"
	fmt.Println(name)
	count:=0
	for index,value := range name{
		fmt.Println(index,value)
		count++
	}
	fmt.Println(count)
}