package main

import "fmt"

// interface definition
type PaymentGetway interface {
	Pay(amount float64) (string, error)
	GetName() string
}

// struct implementing interface (BCA)
type BCAPayment struct{}

func (b BCAPayment) Pay(amount float64) (string, error) {
	return fmt.Sprintf("BCA :pembayaran %.2f berhasil", amount), nil
}
func (b BCAPayment) GetName() string {
	return "Bank BCA"
}

// struct implementing OVO
type OVOPayment struct{}

func (o OVOPayment) Pay(amount float64) (string, error) {
	return fmt.Sprintf("OVO: pembayaran %.2f berhasil", amount), nil
}

func (o OVOPayment) GetName() string {
	return "OVO"
}

// function yang benerima interface
func ProcessPayment(getway PaymentGetway, amount float64) {
	status, _ := getway.Pay(amount)
	fmt.Println(status)
}

func main() {
	var payment PaymentGetway // deklarasi interface

	// bisa ganti implementasi tampa ubah logic core
	payment = BCAPayment{}
	ProcessPayment(payment, 100000)

	payment = OVOPayment{}
	ProcessPayment(payment, 50000)

	payment.GetName()

}
