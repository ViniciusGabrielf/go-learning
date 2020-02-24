package main

import "fmt"

func main() {
	slice1 := make([]int, 10, 20)
	slice2 := append(slice1, 1, 2, 3) // point to same internal array

	fmt.Println(slice1, slice2)

	slice1[0] = 7
	fmt.Println(slice1, slice2)
}
