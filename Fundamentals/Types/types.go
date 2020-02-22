package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer number
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal integer is", reflect.TypeOf(32000))

	// no signal (only positives)... uint8 (byte) uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Byte is", reflect.TypeOf(b))

	// with signal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The max value of int is", i1)
	fmt.Println("Type of i1 is", reflect.TypeOf(i1))

	var i2 rune = 'a' // represent one map of Unicode table (int32)
	fmt.Println("The rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// real numbers (float32, float64)
	var x float32 = 49.99 // var x = float32(49.99)
	fmt.Println("Type of x is", reflect.TypeOf(x))
	fmt.Println("Type of literal 49.99 is", reflect.TypeOf(49.99)) //float64

	// boolean
	bo := true
	fmt.Println("Type of bo is", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hello, my name is Vinicius"
	fmt.Println(s1 + "!")
	fmt.Println("The length of string is", len(s1))

	// string with multiple lines
	s2 := `Hello
	my
	name
	is
	Vinicius`
	fmt.Println(s2 + "!")
	fmt.Println("The length of string is", len(s2))

	// char?? - doens't exist
	char := 'a' // Unicode
	fmt.Println("Type of char is", reflect.TypeOf(char))
	fmt.Println(char)
}
