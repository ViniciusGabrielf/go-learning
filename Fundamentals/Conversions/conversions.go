package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// caution...
	fmt.Println("Teste " + string(97)) // string() converto to ASCII

	// int for string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string for int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}
}
