package main

import (
	"fmt"
	"time"
)

func typeVar(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Floating"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "I don't know"
	}
}

func main() {
	fmt.Println(typeVar(2.3))
	fmt.Println(typeVar(1))
	fmt.Println(typeVar("Ops"))
	fmt.Println(typeVar(func() {}))
	fmt.Println(typeVar(time.Now()))
}
