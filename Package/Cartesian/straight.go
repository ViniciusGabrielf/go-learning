package main

import "math"

// Initializing with a UPPERCASE letter is PUBLIC (visible inside and outside of the package)!
// Initializing with a lowercase letter is PRIVATE (visible only inside the package)!

// Example
// Point - public (public structure need a comment)
// point or _Pointo - private

// Point represent a coordinate in the cartesian plane
type Point struct {
	x float64
	y float64
}

func cathetus(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distance is responsible to calculate linear distance between two points
func Distance(a, b Point) float64 {
	cx, cy := cathetus(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
