package main

import "fmt"

func noteForConcept(value float64) string {
	if value >= 9 && value <= 10 {
		return "Concept B"
	} else if value >= 8 && value < 9 {
		return "Concept B"
	} else if value >= 5 && value < 8 {
		return "Concept C"
	} else if value >= 3 && value < 5 {
		return "Concept D"
	} else {
		return "Concept E"
	}
}

func main() {
	fmt.Println(noteForConcept(9.8))
	fmt.Println(noteForConcept(6.9))
	fmt.Println(noteForConcept(2.1))
}
