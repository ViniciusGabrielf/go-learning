package main

import "fmt"

func printResult(value float64) {
	if value >= 7 {
		fmt.Println("Approved with value", value)
	} else {
		fmt.Println("Reproved with value", value)
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
}
