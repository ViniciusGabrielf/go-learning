package main

import "fmt"

func main() {
	funcAndSalary := map[string]float64{
		"Vinicius Gabriel": 11235.45,
		"André Vaccari":    12331.78,
		"Gabriel Reis":     1200.0,
	}

	funcAndSalary["Vinicius Amaral"] = 4560.0
	delete(funcAndSalary, "do not exist") // no problem

	for name, salary := range funcAndSalary {
		fmt.Println(name, salary)
	}
}
