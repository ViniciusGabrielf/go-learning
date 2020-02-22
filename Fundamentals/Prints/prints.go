package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line.")

	fmt.Println(" New")
	fmt.Println("line.")

	x := 3.141516

	// fmt.Println("Value of x is " + x)

	xs := fmt.Sprint(x)
	fmt.Println("Value of x is" + xs)
	fmt.Println("Value of x is", x)

	fmt.Printf("Value of x is %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
