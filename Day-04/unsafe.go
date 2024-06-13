package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	name string
	id   int
}

func main() {
	w := []int{1, 2, 3, 4}
	ptr := &w[0]
	nextptr := unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(w[0]))
	fmt.Println(*((*int)(nextptr)))
	student1 := Student{"ndv", 1}
	//student2 := Student{"nyhal",2}
	//fmt.Printf("align %d\n", unsafe.Offsetof(student1))
	fmt.Printf("align %d\n", unsafe.Offsetof(student1.name))
	fmt.Printf("align %d\n", unsafe.Offsetof(student1.id))

}
