package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome Home")
}

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(Home))

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Starting server on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
