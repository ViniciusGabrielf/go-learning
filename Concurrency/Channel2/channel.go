package main

import (
	"fmt"
	"time"
)

// Channel - is the communication between goroutines
// Chanel is a type

func twoThreeFourTimes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // send data to channel

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go twoThreeFourTimes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // receive data from channel
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)

	fmt.Println(<-c) // deadlock, all goroutines sleep
}
