package main

//create  different arrays or slices of different data types , check if slice is equal using deepequal in reflection
//sizeof any array
//find the number of elements of array
//reverse the array
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	int_slice1 := []int{1, 2, 3, 4, 5}
	int_slice2 := []int{1, 2, 3, 4, 5}
	string_slice := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("int_slice1 equal string_slice2", reflect.DeepEqual(int_slice1, string_slice))
	fmt.Println("int_slice1 equal int_slice2", reflect.DeepEqual(int_slice1, int_slice2))

	fmt.Println("sizeof int_slice", unsafe.Sizeof(int_slice1))
	fmt.Println("sizeof string slice", unsafe.Sizeof(string_slice))

	fmt.Println(reflect.ValueOf(int_slice1).Len())
	slice1 := []int{1, 2, 3, 4, 5}

	swap := reflect.Swapper(slice1)

	for i := 0; i < len(slice1)/2; i++ {
		swap(i, len(slice1)-i-1)
	}

	fmt.Println(slice1)

}
