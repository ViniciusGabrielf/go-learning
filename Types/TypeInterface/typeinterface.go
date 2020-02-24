package main

import (
	"fmt"
	"reflect"
)

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing, reflect.TypeOf(thing))

	type dinamic interface{}
	var thing2 dinamic = "Opa"
	fmt.Println(thing2, reflect.TypeOf(thing2))

	thing2 = true
	fmt.Println(thing2, reflect.TypeOf(thing2))

	thing2 = course{"Golang: explore the google language."}
	fmt.Println(thing2, reflect.TypeOf(thing2))
}
