package main

import (
	"fmt"

	"github.com/viniciusgabrielf/area"
)

func main() {
	fmt.Println(area.Circle(6.0))
	fmt.Println(area.Rectangle(5.0, 2.0))
	// fmt.Println(area._TriangleEq(5.0,2.0)) => Private function, can not use
}

// CAUTION: This example need a package and file area in github.com/viniciusgabrielf/area (in /src inside GOPATH)
// Code of area.go below:

/* package area

import "math"

// Pi is a numerical proportion defined by the perimeter of a circle and your diameter
const Pi = 3.1416

// Circle is responsible to calculate area of an circle
func Circle(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

// Rectangle is responsible to calculate area of an rectangle
func Rectangle(base, height float64) float64 {
	return base * height
}

// It is not visible (private function)!
func _TriangleEq(base, height float64) float64 {
	return (base * height) / 2
} */
