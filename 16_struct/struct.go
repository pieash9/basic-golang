package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precession
	customer
}

// func newOrder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	// newCustomer := customer{
	// 	name:  "jhone",
	// 	phone: "01627",
	// }

	newOrder := order{
		id:     "1",
		amount: 30,
		status: "confirmed",
		customer: customer{
			name:  "a",
			phone: "555",
		},
	}

	newOrder.customer.name = "pieash"
	fmt.Println(newOrder)
	// myOrder := newOrder("1", 44, "confirmed")

	// fmt.Println(myOrder)

	// if you don't set any value it will be zero
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")

	// fmt.Println(myOrder)

	// myOrder.createdAt = time.Now()

	// fmt.Println(myOrder.status)
	// fmt.Println(myOrder)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(language)

}
