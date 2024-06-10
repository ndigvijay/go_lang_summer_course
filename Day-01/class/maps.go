package main

import "fmt";

func main(){
	// fmt.Println("hello world")
	//map
		// stores values in key value pairs

	marks:=map[string]int{"golang":80,"java":90,"python":100}
	fmt.Println(marks)
	// for i := 0; i < len(marks); i++ {
		// 	fmt.Println(marks[i])
		// }
		for key,value := range marks{
			fmt.Println(key,value)
		}
		marks["python"]=69
		marks["chem"]=70
		fmt.Println(marks["python"])
		fmt.Println(marks)
	}