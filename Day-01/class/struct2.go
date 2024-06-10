package main

import "fmt"

func main(){
	type Person struct{
		data int
		id int
	}
	Person1:=Person{10,20}
	// fmt.Println(Person1)
	Person2:=Person{id:10,data:20}
	fmt.Println(Person1,Person2)
	fmt.Println(Person2.id,Person2.data)


}