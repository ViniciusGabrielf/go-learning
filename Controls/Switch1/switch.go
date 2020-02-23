package main

import "fmt"

func noteForConcept(n float64) string {
	var note = int(n)
	switch note {
	case 10:
		fallthrough // no break case
	case 9:
		return "A"
	case 8, 7: // test two values
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: // no enter in any cases
		return "Invalid note"
	}
}

func main() {
	fmt.Println(noteForConcept(9.8))
	fmt.Println(noteForConcept(6.9))
	fmt.Println(noteForConcept(2.1))
	fmt.Println(noteForConcept(11.1))
	fmt.Println(noteForConcept(-1.1))
}
