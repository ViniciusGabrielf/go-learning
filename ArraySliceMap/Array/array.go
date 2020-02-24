package main

import "fmt"

func main() {
	// same type of data and static (fixed length)
	var notes [3]float64
	fmt.Println(notes)

	notes[0], notes[1], notes[2] = 7.8, 4.3, 9.1
	// notes[3] = 10 // invalid index
	fmt.Println(notes)

	total := 0.0
	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	average := total / float64(len(notes))
	fmt.Printf("Average %.2f\n", average)
}
