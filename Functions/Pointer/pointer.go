package main

import "fmt"

func inc1(n int) {
	n++
}

// pointer is represented by *
func inc2(n *int) {
	// * in this case is to resolve pointer => access value of pointer reference
	*n++
}

func main() {
	n := 1

	inc1(n) // per value
	fmt.Println(n)

	// & get address from an variable
	inc2(&n)
	fmt.Println(n)
}
