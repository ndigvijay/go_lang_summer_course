package main

import "fmt"

type shape interface{
	area() float64
	perimeter() float64
}

type square struct{
	length float64
}

type circle struct{
	radius float64
}


func(s square) area()float64{
	return s.length*s.length
}

func(s square) perimeter()float64{
	return 4*s.length
}

func(c circle) perimeter()float64{
	return 2*3.14*c.radius
}


func(c circle) area()float64{
	return 3.14*c.radius*c.radius
}

func PrintShapeInfo(s shape){
	fmt.Printf("area of %v is %f",s,s.area())
	fmt.Printf("area of %v is %f",s,s.perimeter())

}

func main(){
	// fmt.Println("hello world")
	shapes:=[]shape{
		square{length:15.12},
		circle{radius:7.5},
		circle{radius:12.3},
		square{length:4.9},
	}

	for _,v:=range shapes{
		PrintShapeInfo(v)
		fmt.Println("----")
	}
}