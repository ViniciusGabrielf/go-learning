package main

import (
	"fmt"
	"time"
)

// Multiplexing pattern with generator pattern

func speak(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking: %d", person, i)
		}
	}()
	return c
}

func join(enter1, enter2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-enter1:
				c <- s
			case s := <-enter2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := join(speak("Vinicius"), speak("João"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
