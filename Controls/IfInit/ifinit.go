package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := randomNumber(); i > 5 {
		fmt.Println("Win!!!")
	} else {
		fmt.Println("Lose!!!")
	}

	// fmt.Println(i) -> it's not possible, because i is in for scope
}
