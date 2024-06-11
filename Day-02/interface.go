package main 

import "fmt"

func main() {
	var i interface{} = "Hi"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	y, ok := i.(float32)
	fmt.Println(y, ok)
}