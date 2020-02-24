package main

import "fmt"

func main() {
	funcPerLetter := map[string]map[string]float64{
		"G": {
			"Gabriel Reis":  15564.0,
			"Gabriel Silva": 45453.78,
		},
		"V": {
			"Vinicius Amaral": 12314.45,
		},
		"A": {
			"Andre Vaccari": 1200.0,
		},
	}

	delete(funcPerLetter, "V")

	for letter, funcs := range funcPerLetter {
		fmt.Println(letter, funcs)
	}
}
