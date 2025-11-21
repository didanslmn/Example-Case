package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	// tabel driven test
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"kelipatan 3", 3, "Fizz"},
		{"kelipatan 5", 5, "Buzz"},
		{"kelipatan 15", 15, "FizzBuzz"},
		{"bukan kelipatan", 7, "7"},
		{"nol", 0, "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FizzBuzz(tt.input)
			if result != tt.expected {
				t.Errorf("fizzBuzz(%d) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

// benchmark test
func BenchMarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(100)
	}

}
