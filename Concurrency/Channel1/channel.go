package main

import "fmt"

func main() {
	// channel is a type in go
	ch := make(chan int, 1)

	ch <- 1 // sending data to a channel (write)
	<-ch    // receive data from chanel (read)

	ch <- 2
	fmt.Println(<-ch)
}
