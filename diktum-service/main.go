package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanri/diktum-service/handlers"
	"github.com/lanri/diktum-service/middleware"
)

func main() {
	r := mux.NewRouter()

	fmt.Println("Server berjalan di http://localhost:8083")

	// OAuth2 routes
	r.HandleFunc("/oauth2/login", handlers.OAuthLogin).Methods("GET")
	r.HandleFunc("/oauth2/callback", handlers.OAuthCallback).Methods("GET")

	// Protected route (JWT)
	r.Handle("/api/diktum/rumusan", middleware.AuthMiddleware(http.HandlerFunc(handlers.HelloDiktum))).Methods("GET")

	http.ListenAndServe(":8083", r)
}
