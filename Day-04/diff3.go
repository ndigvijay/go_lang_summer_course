package main

import (
	"fmt"
	"reflect"
)

type Company struct {
	name      string
	location  string
	employees []Employee
}

type Employee struct {
	name        string
	id          int
	designation string
}

func main() {
	appleCompany := Company{
		name:     "Apple",
		location: "Bangalore",
		employees: []Employee{
			{
				name:        "Nihal",
				id:          1,
				designation: "Frontend Developer",
			},
			{
				name:        "anirudh",
				id:          2,
				designation: "figma expert",
			},
		},
	}

	fmt.Println(reflect.ValueOf(appleCompany.employees[0]))
	fmt.Println(reflect.ValueOf(appleCompany.employees[1]))
	fmt.Println(reflect.ValueOf(appleCompany.location))
	fmt.Println(reflect.DeepEqual(appleCompany.location, appleCompany.name))

}
