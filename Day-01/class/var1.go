package main

import "fmt";

func main(){
	// var variable_name datatype=value
	// var nyhal int=69
	// fmt.Println(nyhal)
	// digvijay:="digvijay"
	// fmt.Println(digvijay)
	// var dmello="dmells"
	// fmt.Println(dmello)
	var a int; //0
	var b bool; //false
	var c string;
	var d float32; //o
	fmt.Println(a,b,c,d)

	//declaring multiple variables in a single line
	var rr,ec,hn string="ring road","electronic city","hanumantha nagar"
	fmt.Println(rr,ec,hn)

	// cannot change the datatype of a go lang
	//ERROR if u try to change the type

	//example:
	car:="mercedes"
	// car="redbull"
	// car=10
	fmt.Println(car)
}

