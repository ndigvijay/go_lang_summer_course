package main

import (
	"fmt"
	"strconv"
)

// strconv package - str to int and int to str
func main() {
	s := "hello"
	fmt.Println(len(s))
	fmt.Println(s[0], s[4]) //ascii or unicode
	fmt.Println(s[0:5])     //entire string
	//concatination
	fmt.Println(s[1:] + "nyhal")
	x := "123"
	var y, err = strconv.Atoi(x)
	fmt.Println(y, err)
	z := strconv.Itoa(y)
	fmt.Println(z)
	z1 := "123.45"
	fmt.Println(z1)

}
