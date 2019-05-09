package main

import (
	"fmt"
)

func generateFibonacci(x int) []int64 {
	sequence := []int64{0,1}

	for i := 2; i < x; i++ {
		sequence = append(sequence, sequence[i-1] + sequence[i-2])
	}

	return sequence
}

func main() {
	var sequence = generateFibonacci(93)
	for k, e := range sequence {
		fmt.Printf("%d, %d\n", k, e)
	}
}