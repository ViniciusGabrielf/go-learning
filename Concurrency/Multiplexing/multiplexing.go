package main

import (
	"fmt"

	"github.com/viniciusgabrielf/html"
)

// Multiplexing is a pattern for work with concurrency

func forward(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin
	}
}

// multiplexing - mix (messages) in one channel

func join(enter1, enter2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(enter1, c)
	go forward(enter2, c)
	return c
}

func main() {
	c := join(
		html.Title("https://www.youtube.com.br", "https://www.google.com.br"),
		html.Title("https://www.twitch.tv/", "https://pt.stackoverflow.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
