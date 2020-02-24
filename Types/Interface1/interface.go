package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

// interfaces are implemented implicity
func (p person) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{"Roberto", "Silva"}
	fmt.Println(thing.toString())
	print(thing)

	thing = product{"Calça Jeans", 79.90}
	fmt.Println(thing.toString())
	print(thing)

	p2 := product{"Calça Jeans", 179.90}
	print(p2)
}
