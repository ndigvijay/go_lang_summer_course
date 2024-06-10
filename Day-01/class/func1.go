package main

import "fmt";

func sayhello(name string){
	fmt.Println("hello ",name)
}

func names(s[] string,function func(string)){
	for _,v:=range s{
		sayhello(v)
	}
}

func main(){
	// sayhello("ndv")
	names([]string{"nihal","ndv","nishanthdmello"},sayhello)
}
