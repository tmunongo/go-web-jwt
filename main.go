package main

import (
	"log"
	"net/http"

	"github.com/tmunongo/go-web-jwt/handlers"
)

func main () {
	http.HandleFunc("/", handlers.HomeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
