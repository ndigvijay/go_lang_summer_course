package main

import (
	"fmt"
	"reflect"
)

func main() {

	slices := []string{"go programming", "go lang"}

	copySlice := make([]string, len(slices))

	src := reflect.ValueOf(slices)
	dst := reflect.ValueOf(copySlice)
	reflect.Copy(dst, src)

	fmt.Println("Original slice:", slices)
	fmt.Println("Copied slice:", copySlice)
}
