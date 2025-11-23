package main

import (
	"fmt"
	"log"
)

func pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error : pembagi tidak boleh nol")
	}
	return a / b, nil
}

func main() {
	hasil, err := pembagian(10, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hasil pembagian adalah", hasil)
}
