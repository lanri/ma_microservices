package handlers

import (
	"fmt"
	"net/http"
)

func HelloDiktum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Diktum Service!")
}
