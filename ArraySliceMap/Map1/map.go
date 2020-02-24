package main

import "fmt"

func main() {
	// var approved map[int]string
	// map should be initialized
	approved := make(map[int]string)

	approved[1234567878] = "Maria"
	approved[9823982101] = "Pedro"
	approved[1293123900] = "Ana"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s {CPF: %d}\n", name, cpf)
	}

	fmt.Println(approved[9823982101])
	delete(approved, 9823982101) // delete position from map, by index
	fmt.Println(approved)
	fmt.Println(approved[9823982101])
}
