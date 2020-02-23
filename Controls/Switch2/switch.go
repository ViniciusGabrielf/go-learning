package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 10:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good nigh!")
	}
}
