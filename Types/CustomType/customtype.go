package main

import "fmt"

type value float64

func (v value) between(init, finish float64) bool {
	return float64(v) >= init && float64(v) <= finish
}

func valueForConcept(v value) string {
	if v.between(9.0, 10.0) {
		return "A"
	} else if v.between(7.0, 8.99) {

	} else if v.between(5.0, 7.99) {
		return "C"
	} else if v.between(3.0, 4.99) {
		return "D"
	}
	return "E"
}

func main() {
	fmt.Println(valueForConcept(9.8))
	fmt.Println(valueForConcept(6.9))
	fmt.Println(valueForConcept(2.1))
}
