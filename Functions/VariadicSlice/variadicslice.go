package main

import "fmt"

func printApproved(approved ...string) {
	fmt.Println("List of Approved")
	for i, student := range approved {
		fmt.Printf("%d) %s\n", i+1, student)
	}
}

func main() {
	approved := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	printApproved(approved...)
}
