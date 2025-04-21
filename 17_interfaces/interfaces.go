package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

// open close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	p.gateway.pay(amount)
	// razorpayPaymentGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct {
}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making oatment usign fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal")
}

func (p paypal) refund(amount float32, account string) {
	fmt.Println("Making refund using paypal")
}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	// fakeGw := fakepayment{}
	paypalGw := paypal{}

	newPayment := payment{
		gateway: paypalGw,
	}
	// newPayment := payment{}
	newPayment.makePayment(100)
	newPayment.gateway.refund(100, "123")
}
