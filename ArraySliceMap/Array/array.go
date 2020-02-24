package main

import "fmt"

func main() {
	// same type of data and static (fixed length)
	var values [3]float64
	fmt.Println(values)

	values[0], values[1], values[2] = 7.8, 4.3, 9.1
	// values[3] = 10 // invalid index
	fmt.Println(values)

	total := 0.0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	average := total / float64(len(values))
	fmt.Printf("Average %.2f\n", average)
}
