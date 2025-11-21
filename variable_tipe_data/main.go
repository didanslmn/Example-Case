package main

import "fmt"

func main() {
	var name string = "budi"
	var umur int = 20

	// short variable declaration
	alamat := "Pemalang"
	suhu := 29.8

	const pi = 3.14

	fmt.Printf("Halo nama saya %s umur saya %d tahun\n", name, umur)
	fmt.Printf("alamat %s, suhu %f\n", alamat, suhu)
	fmt.Printf("pi %f", pi)

}
