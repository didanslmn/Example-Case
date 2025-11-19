package main

import "fmt"

func kalkulatorVariadic(numbers ...int) (sum int, count int) {
	count = len(numbers)
	for _, num := range numbers {
		sum += num
	}
	return
}

func main() {
	sum, count := kalkulatorVariadic(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d, Count: %d\n", sum, count)
}
