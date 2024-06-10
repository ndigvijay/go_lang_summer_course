package main

import "fmt"

func main(){
	fmt.Println("hello world")
	s:=[]int{10,20,30,40,50}
	fmt.Println(s)
	fmt.Println(s[1:4])// start to end-1
	slice1:=make([]int,3,5) //length=3 capacity=5
	fmt.Println(slice1)
	//len(),//cap()
	slice1=append(slice1,4,5,6)
	fmt.Println(slice1)

}