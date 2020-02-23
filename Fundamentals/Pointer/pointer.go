package main

import "fmt"

func main() {
	i := 1

	var pointer *int = nil

	pointer = &i // get address from variable i
	*pointer++   // resolve address / get value
	i++          // use same value in *pointer

	// Go hasn't pointer arithmetic
	// pointer++

	fmt.Println(pointer, *pointer, i, &i) // &i get address from variable i
}
