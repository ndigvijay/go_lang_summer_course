package main

import "fmt"

func main()  {
	fmt.Println("sike")
	// no args no return value 
	fmt.Println(sub(10,5))
	fmt.Println(pairs(10,5))
}

func add(a int,b int){
	//2 args no return value
}

func sub(a int,b int)int{
	// 2 args return int
	return a-b
}

func pairs(a int ,b int)(int,int){
	return a,b
}