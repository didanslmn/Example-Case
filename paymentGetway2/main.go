package main

import "fmt"

// PaymentGateway defines the interface for payment gateways
type PaymentGateway interface {
	Pay(amount float64) (string, error)
	GetName() string
}

// CreditCardPayment implements the PaymentGateway interface for credit card payments
type CreditCard struct {
	HolderName string
	Number     string
}

func (cc CreditCard) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Credit Card: pembayaran %.2f dibayar oleh %s", amount, cc.HolderName), nil
}

func (cc CreditCard) GetName() string {
	return "Credit Card"
}

// PayPalPayment implements the PaymentGateway interface for PayPal payments
type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) (string, error) {
	return fmt.Sprintf("PayPal: pembayaran %.2f dibayar oleh %s", amount, pp.Email), nil
}
func (pp PayPal) GetName() string {
	return "PayPal"
}

// function yang menerima interface (polymorphism)
func OrderProcessing(getway PaymentGateway, amount float64) {
	fmt.Printf("Menggunakan %s untuk membayar\n", getway.GetName())

	result, err := getway.Pay(amount)
	if err != nil {
		fmt.Printf(" error: %v\n", err)
		return
	}
	fmt.Println(result)
	fmt.Println("order selesai")
}

func main() {

	cc := CreditCard{
		HolderName: "Didan",
		Number:     "1234-5678-9012-3456",
	}
	pp := PayPal{
		Email: "william.johnson@example-pet-store.com",
	}
	OrderProcessing(pp, 150000)
	OrderProcessing(cc, 200000)
}
