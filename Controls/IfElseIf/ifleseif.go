package main

import "fmt"

func noteForConcept(note float64) string {
	if note >= 9 && note <= 10 {
		return "Concept B"
	} else if note >= 8 && note < 9 {
		return "Concept B"
	} else if note >= 5 && note < 8 {
		return "Concept C"
	} else if note >= 3 && note < 5 {
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
