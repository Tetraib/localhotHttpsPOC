package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Localhost called")
	})
	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://foo.com"},
		// AllowedMethods: []string{"0"},
	})
	handler := cors.Default().Handler(mux)
	handler = c.Handler(handler)
	http.ListenAndServe(":8080", handler)
}
