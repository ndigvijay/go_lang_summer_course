package main

import "fmt";
//user defined datatypes
// syntax

// type struct_name struct{
		// members datatype
// }

type Person struct{
	name string
	age int
}


func main()  {
	//struct variables
	var p1 Person
	var p2 =Person{name:"N Digvijay",age:20}
	p3:=Person{name: "nynynyhal",age:45}
	p4:=&Person{name: "nynynyhal",age:45}
	p5:=struct{
		data int
	}{data:40};//anon struct
	fmt.Println(p1,p2,p3,p4,p5)
}