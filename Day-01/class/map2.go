package main

import "fmt";

func main(){
	fmt.Println("hello world")
	person:=map[string]int{"alice":30,"ndv":20}
	fmt.Println(person)
	delete(person,"alice")
	fmt.Println(person)
}