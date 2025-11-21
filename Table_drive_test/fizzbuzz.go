package main

import "fmt"

func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "fizz buzz"
	} else if n%3 == 0 {
		return "fizz"
	} else if n%5 == 0 {
		return "buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}
