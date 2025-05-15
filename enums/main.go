package main

import "fmt"

// enumrated types
type OrderStatus string

const (
	Received   OrderStatus = "received"
	Processing             = "processing"
	Shipped                = "shipped"
	Delivered              = "delivered"
	Cancelled              = "cancelled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Received)
}
