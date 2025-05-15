package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{ gateway paymenter }

type razorpay struct{}

type stripe struct{}

type paypal struct{}

type fakePayment struct{}

func (p payment) makePayment(amount float32) {
	// implementation
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)
	// stripePaymentGW := stripe{}
	p.gateway.pay(amount)
}

func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay payment of amount:", amount)
}

func (s stripe) pay(amount float32) {
	fmt.Println("Stripe payment of amount:", amount)
}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Fake payment of amount:", amount)
}

func (f paypal) pay(amount float32) {
	fmt.Println("Paypal payment of amount:", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("Razorpay refund of amount:", amount, "to account:", account)
}
func (s stripe) refund(amount float32, account string) {
	fmt.Println("Stripe refund of amount:", amount, "to account:", account)
}
func (f fakePayment) refund(amount float32, account string) {
	fmt.Println("Fake payment refund of amount:", amount, "to account:", account)
}
func (f paypal) refund(amount float32, account string) {
	fmt.Println("Paypal refund of amount:", amount, "to account:", account)
}

func main() {
	stripePaymentGW := stripe{}
	razorpayPaymentGW := razorpay{}
	paypalPaymentGW := paypal{}
	fakePaymentGW := fakePayment{}

	Payment1 := payment{
		gateway: razorpayPaymentGW,
	}
	Payment2 := payment{
		gateway: stripePaymentGW,
	}
	Payment3 := payment{
		gateway: paypalPaymentGW,
	}
	newPayment := payment{
		gateway: fakePaymentGW,
	}

	Payment1.makePayment(100.50)
	Payment2.makePayment(200.50)
	Payment3.makePayment(300.50)
	newPayment.makePayment(400.50)

	// make refund
	Payment1.gateway.refund(100.50, "1234567890")
	Payment2.gateway.refund(200.50, "1234567890")
	Payment3.gateway.refund(300.50, "1234567890")
	newPayment.gateway.refund(400.50, "1234567890")
}

//  Open close principle
// 1. Open for extension
// 2. Closed for modification
// 3. New functionality can be added without modifying existing code
// 4. New payment gateways can be added without modifying existing code
// 5. New payment gateways can be added by implementing the pay method
