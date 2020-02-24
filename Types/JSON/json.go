package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := product{1, "Notebook", 1899.90, []string{"Sale", "Eletronic"}}
	p1Json, _ := json.Marshal(p1) // Convert struct to array of bytes
	// fmt.Println(p1Json) => in bytes
	fmt.Println(string(p1Json))

	// json to struct
	var p2 product
	jsonString := `{"id":2, "name":"Caneta","price":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2) // Receive an array of bytes to convert and insert into a interface
	fmt.Println(p2.Tags[1])
}
