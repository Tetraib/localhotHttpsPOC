package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://eager-boyd-0fa197.netlify.com"},
	})
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Localhost called")
	})
	http.ListenAndServe(":8080", c.Handler(handler))
}
