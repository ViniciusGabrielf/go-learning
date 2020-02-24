package main

import "fmt"

type sport interface {
	turnOnTurbo()
}

type ferrari struct {
	model        string
	turboOn      bool
	currentSpeed int
}

func (f *ferrari) turnOnTurbo() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnTurbo()

	var car2 sport = &ferrari{"F40", false, 0}

	fmt.Println(car1, car2)
}
