package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car     // anonymous fields (composition)
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = true

	fmt.Printf("Ferrari %s turbo is on? %v\n", f.name, f.turboOn)
	fmt.Println(f)
}
