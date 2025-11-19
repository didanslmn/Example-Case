package main

import "fmt"

func swap(a, b *int) {
	// swap dengan pointer
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 5, 10
	fmt.Println("sebelum", x, y)
	swap(&x, &y)
	fmt.Println("setelah", x, y)
}
