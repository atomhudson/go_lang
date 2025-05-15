package main

import (
	"fmt"
	"time"
)

// order struct

type Customer struct {
	name  string
	phone string
}

type Order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
	Customer
}

func newOrder(id int, amount float32, status string) *Order {

	myOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}

func (o *Order) changeOrderStatus(status string) {
	o.status = status
}
func (o *Order) getOrderStatus() string {
	return o.status
}
func (o *Order) getOrderAmount() float32 {
	return o.amount
}
func (o *Order) getOrderCreatedAt() time.Time {
	return o.createdAt
}
func (o *Order) getOrderID() int {
	return o.id
}

func main() {
	order := Order{
		id:        1,
		amount:    100.50,
		status:    "pending",
		createdAt: time.Now(),
	}
	fmt.Println("Order :", order)
	fmt.Println("==========================")
	order.changeOrderStatus("completed")
	fmt.Println("Order ID:", order.id)
	fmt.Println("Order Amount:", order.amount)
	fmt.Println("Order Status:", order.status)
	fmt.Println("Order Created At:", order.createdAt)
	fmt.Println("==========================")
	fmt.Println("Order ID:", order.getOrderID())
	fmt.Println("Order Amount:", order.getOrderAmount())
	fmt.Println("Order Status:", order.getOrderStatus())
	fmt.Println("Order Created At:", order.getOrderCreatedAt())
	fmt.Println("==========================")
	myOrder := newOrder(2, 200.75, "pending")
	fmt.Println("Order: ", myOrder)
	// Print the order details
	fmt.Println("==========================")
	language := struct {
		name    string
		version string
	}{"Go", "1.20"}
	fmt.Println("Language Struct: ", language)

	fmt.Println("==========================")

	newOrder := Order{
		id:        3,
		amount:    150.00,
		status:    "pending",
		createdAt: time.Now(),
		Customer: Customer{
			name:  "John Doe",
			phone: "123-456-7890",
		},
	}

	fmt.Println("New Order :", newOrder)

}
