package main

import "fmt"


func main(){
	fmt.Println("hello world")
	// var array_name=[length]datatype{values}length is fixed
	// var array_name=[...]datatype{values}length is taken implicitly
	//walrus the goat operator:
	// array_name:=[length]datatype{values}

	k:=[...]int{99:-1}
	fmt.Println(k)
	fmt.Println(k[len(k)-1])

	numbers:=[2]int{1,2}
	values:=[...]int{1,2}
	no_yo:=[2]int{1,3}
	fmt.Println(numbers==values)
	fmt.Println(values==no_yo)
	fmt.Println(no_yo==numbers)
}