package main
import "fmt";
func main()  {
	type Employee struct{
		id int
		name string
		address string
	}

	Employee1:=Employee{1,"digvijay","PESU EC"}
	Employee2:=Employee{id:2,name:"nyhal",address: "PESU RR kempus"}
	Employee3:=Employee{id:3}
	Nishanth_Pointer:=&Employee3.name
	*Nishanth_Pointer="Nishanth Dmello"
	fmt.Println(Employee1,Employee2,Employee3)

	//go doesnt support pointer arthmetic

}