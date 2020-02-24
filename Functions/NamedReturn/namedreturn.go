package main

import "fmt"

func exchange(p1, p2 int) (second int, first int) {
	second = p2
	first = p1
	return // clean return
}

func main() {
	x, y := exchange(2, 3)
	fmt.Println(x, y)

	second, first := exchange(7, 1)
	fmt.Println(second, first)
}
