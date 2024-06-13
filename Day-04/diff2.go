package main

import (
	"fmt"
	"unsafe"
)

//create a struct pes , consisting of courses rank location
//access the structure memebers
//find the size of sturcture
// perform align and offset of struct names

type PES struct {
	courses []string
	rank    int
}

func main() {
	PESU_EC := PES{[]string{"cse", "ece"}, 2}
	PESU_RR := PES{[]string{"cse", "ece", "mech", "psy", "civil", "eee", "bio"}, 1}
	fmt.Println(PESU_EC, PESU_RR)
	fmt.Println(unsafe.Sizeof(PESU_EC.courses), unsafe.Sizeof(PESU_RR.courses))
	fmt.Println(unsafe.Alignof(PESU_EC.courses), unsafe.Alignof(PESU_RR.courses))
	fmt.Println(unsafe.Offsetof(PESU_EC.courses), unsafe.Sizeof(PESU_RR.courses))
}
