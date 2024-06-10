package main

import "fmt"

type User struct{
	name string
	age int
}
// Method is 

func main(){
	var u User
	u.name="ndv"
	u.age=21
	name:=u.full_name()
	fmt.Println(name)
}

//method
// attach structure with function
func(u User) full_name()string{
	return fmt.Sprintf("%s age %d",u.name,u.age)
}