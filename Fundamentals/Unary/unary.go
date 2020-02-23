package main

import "fmt"

func main() {
	x := 1
	y := 2

	// go has only postfix form
	x++ // x += 1 or x = x +1
	fmt.Println(x)

	y-- // x -= 1 or y = y - 1
	fmt.Println(y)
}
