package main

import "fmt"

type sport interface {
	turnOnTurbo()
}

type luxurious interface {
	makeParallelPark()
}

type sportLuxurious interface {
	sport
	luxurious
	// can add new methods
}

type bmw7 struct{}

func (b bmw7) turnOnTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) makeParallelPark() {
	fmt.Println("Parallel Park...")
}

func main() {
	var b sportLuxurious = bmw7{}
	b.turnOnTurbo()
	b.makeParallelPark()

}
