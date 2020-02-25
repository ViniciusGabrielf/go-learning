package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executed 1!")
	ch <- 4 // script locked by buffer, max length (3)
	fmt.Println("Executed 2!")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
