package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Generator is a pattern for work with concurrency
// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal only-read
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, _ := http.Get(url)
			html, _ := ioutil.ReadAll(res.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	// Concurrency encapsulated in function
	// Function call don't see concurrency
	t1 := title("https://www.google.com.br", "https://www.twitch.tv/")
	t2 := title("https://pt.stackoverflow.com", "https://www.youtube.com")
	fmt.Println("First: ", <-t1, "|", <-t2)
	fmt.Println("Second: ", <-t1, "|", <-t2)
}
