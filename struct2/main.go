package main

import "fmt"

type Product struct {
	ID          string
	Name        string
	Price       float64
	Stock       int
	isAvailabel bool
}

func newProduct(id, name string, price float64) Product {
	return Product{
		ID:          id,
		Name:        name,
		Price:       price,
		Stock:       0,
		isAvailabel: false,
	}
}

// cek ketersediaan produck

func (p *Product) IsReady() bool {
	return p.isAvailabel && p.Stock > 0

}

func main() {
	product1 := newProduct("1", "laptop", 15000000)
	product1.Stock = 10
	product1.isAvailabel = true

	// akses field

	fmt.Println("nama product:", product1.Name)
	fmt.Printf("tersedia?: %v\n", product1.IsReady())

}
