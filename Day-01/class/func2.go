package main

import "fmt"

func main(){
	var x int =5
	fmt.Println(x)
	fmt.Println(f1(&x))
	fmt.Println(x)
}

func f1(number *int)int{
	*number++
	return *number
}
