package main

import "fmt"

func printResult(note float64) {
	if note >= 7 {
		fmt.Println("Approved with note", note)
	} else {
		fmt.Println("Reproved with note", note)
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
}
