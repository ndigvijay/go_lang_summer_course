package main

import (
	"fmt";
)

func main(){
	fmt.Println("hello world")
	slice1:=make([]int,3,5)
	for index,value := range slice1{
		fmt.Println(index,value)
	}
	//method 2:
	for i:=0;i<len(slice1);i++{
		fmt.Println(slice1[i])
	}

	// appending slices
	slice69:=[]int{1,2,3,45};
	slice70:=[]int{20,21,22};
	slice420:=append(slice69,slice70...)
	fmt.Println(slice420)

	var stack[]string
	// top:=-1
	v:="apple"
	// top++
	stack=append(stack,v)
	fmt.Println(stack)
	// stack[top]=v
	// println(stack)

	top:=stack[len(stack)-1]
	fmt.Println(top)
	
	pop:=stack[:len(stack)-1]
	fmt.Println(pop)
}