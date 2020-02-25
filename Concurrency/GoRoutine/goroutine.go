package main

import (
	"fmt"
	"time"
)

func speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// linear execution
	// #TEST 1
	// speak("Vinicius", "Why you don't speak with me?", 3)
	// speak("João", "I can speak only after you!", 1)

	// reserved word "go" create a goroutine for that function (open a independent execution)
	// only execute while the main thread of the pogram is executing

	// #TEST 2
	// go speak("Vinicius", "Ei...", 500)
	// go speak("João", "Opa...", 500)

	// time.Sleep(time.Second * 5)
	// fmt.Println("Finish!")

	// #TEST 3
	go speak("Vinicius", "Right!", 10)
	speak("João", "Congratulations!", 5)
}
