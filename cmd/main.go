package main

import (
	"log"
	"net/http"

	"github.com/tmunongo/go-web-jwt/pkg/handlers"
	"github.com/tmunongo/go-web-jwt/pkg/middleware"
)

func main () {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/admin", middleware.AuthMiddleware(handlers.ProtectedHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
