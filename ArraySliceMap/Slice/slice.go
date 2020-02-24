package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [3]int{1, 2, 3} // array (fixed length)
	slice1 := []int{1, 2, 3}  // slice (variable length)

	fmt.Println(array1, slice1)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [5]int{1, 2, 3, 4, 5}

	// Slice is not an array! Slice defines a part of an array
	// Slice has a length and one pointer for one element of an array

	slice2 := array2[1:3]
	fmt.Println(array2, slice2)

	slice3 := array2[:2] // new slice, point to same array
	fmt.Println(array2, slice3)

	slice4 := slice2[:1] // slice from slice
	fmt.Println(slice2, slice4)
}
