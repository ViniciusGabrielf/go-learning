package main

import (
	"fmt"
	"time"

	"github.com/viniciusgabrielf/html"
)

func theFastest(url1, url2, url3 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	// structure control specific for concurrency
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "All lose!"
		// default:
		// 	return "No response!"
	}
}

func main() {
	champion := theFastest(
		"https://www.youtube.com.br",
		"https://www.google.com.br",
		"https://pt.stackoverflow.com",
	)
	fmt.Println(champion)
}
