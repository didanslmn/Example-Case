package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go go lang belajar golang go go go"
	wordCount := make(map[string]int)

	// split string menjadi slice
	words := strings.Fields(text)

	// hitung frekuensi
	for _, word := range words {
		wordCount[word]++
	}

	// print hasil
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

}
