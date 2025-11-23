package main

import "fmt"

func hitungLuas(panjang, lebar float64) float64 {
	luas := panjang * lebar
	return luas
}

func main() {

	luasPersegiPanjang := hitungLuas(9, 5)
	fmt.Println("Luas persegi panjang adalah:", luasPersegiPanjang)
}
