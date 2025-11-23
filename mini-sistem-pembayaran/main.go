package main

import "fmt"

// struct ( data)
type Invoice struct {
	ID     string
	Amount float64
	Status string
}

// interface (kontrak)
type PaymentMethod interface {
	Pay(invoice Invoice) error
	GetProvider() string
}

// implementasi : cash
type CashPayment struct {
	CashierName string
}

func (c CashPayment) Pay(invoice Invoice) error {
	fmt.Printf("[%s] bayar tunai sebesar %.2f ke %s\n", invoice.ID, invoice.Amount, c.CashierName)
	return nil
}

func (c CashPayment) GetProvider() string {
	return "cash"
}

// ============ IMPLEMENTASI: QRIS ============
type QRISPayment struct {
	MerchantID string
}

func (q QRISPayment) Pay(invoice Invoice) error {
	fmt.Printf("[%s] Scan QRIS $%.2f (Merchant: %s)\n",
		invoice.ID, invoice.Amount, q.MerchantID)
	return nil
}

func (q QRISPayment) GetProvider() string {
	return "QRIS"
}

// ============ MAIN: Simulasi ============
func checkout(invoice Invoice, method PaymentMethod) {
	fmt.Printf("\n🛒 Checkout Invoice %s: $%.2f\n", invoice.ID, invoice.Amount)
	fmt.Printf("💳 Payment: %s\n", method.GetProvider())
	method.Pay(invoice)
	invoice.Status = "PAID" // ini tidak akan mengubah status di main karena pass by value
	fmt.Printf("✅ Invoice %s status updated to %s\n", invoice.ID, invoice.Status)
}

func main() {
	// Buat invoice
	inv := Invoice{
		ID:     "INV-001",
		Amount: 250000,
		Status: "PENDING",
	}

	// Pilih payment method
	cash := CashPayment{CashierName: "Siti"}
	qris := QRISPayment{MerchantID: "M-12345"}

	// Checkout dengan method berbeda
	checkout(inv, cash)
	checkout(inv, qris)

	// Cek status (masih PENDING karena pass by value!)
	fmt.Printf("\nStatus invoice: %s\n", inv.Status)
}
