package main

//platform independant to os functions
// os functions,env variables,command line args etc
import (
	"fmt"
	"os"
);

func main(){
	fmt.Println("hello")
	// args:=os.Args= all arguements will be printed
	args1:=os.Args[1:] // 1 to end
	// args2=os.Args[1]

	fmt.Println(args1)
}
