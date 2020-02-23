package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // has no while in Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMixing structures...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Pair")
		} else {
			fmt.Println("Odd")
		}
	}

	for {
		// infinite loop
		fmt.Println("For ever...")
		time.Sleep(time.Second)
	}
}
