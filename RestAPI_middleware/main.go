package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// middleware logger

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("%s %s - Started\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		fmt.Printf("%s %s - Completed in %v\n", r.Method, r.URL.Path, duration)
	})
}

// middleware authsederhana
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// handler

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	response := map[string]string{"message": "Hello from go api"}
	json.NewEncoder(w).Encode(response)

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	// chain middleware
	handler := loggingMiddleware(authMiddleware(mux))

	fmt.Println("runing server on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
