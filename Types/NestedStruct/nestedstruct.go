package main

import "fmt"

type item struct {
	productID int
	amount    int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (p order) totalPrice() float64 {
	total := 0.0
	for _, item := range p.items {
		total += item.price * float64(item.amount)
	}
	return total
}

func main() {
	order := order{
		userID: 1,
		items: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Total value of order is R$%.2f", order.totalPrice())
}
