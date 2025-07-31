package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lanri/diktum-service/utils"
)

func OAuthLogin(w http.ResponseWriter, r *http.Request) {
	url := utils.GetAuthURL("some-random-state")
	http.Redirect(w, r, url, http.StatusFound)
}

func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	token, err := utils.ExchangeCode(context.Background(), code)
	if err != nil {
		http.Error(w, "OAuth exchange failed", http.StatusInternalServerError)
		return
	}

	// Token dapat digunakan untuk ambil userinfo jika perlu
	fmt.Fprintf(w, "Login berhasil!\nToken: %s\nExpired at: %s", token.AccessToken, token.Expiry)
}
