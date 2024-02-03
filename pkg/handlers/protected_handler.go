package handlers

import "net/http"

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, User!"))
}