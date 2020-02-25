package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // blocking operation
	fmt.Println("Only after chanel is read")
}

func main() {
	c := make(chan int) // channel without buffer

	go routine(c)

	fmt.Println(<-c) // blocking operation
	fmt.Println("Data read")
	fmt.Println(<-c) // deadlock
	fmt.Println("Finish!")
}
