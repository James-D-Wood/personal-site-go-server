package middleware

import (
	"fmt"
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement real auth
		authKey := os.Getenv("PS_Auth_Key")
		if authKey == "" {
			http.Error(w, "Auth Error", http.StatusInternalServerError)
			return
		}
		auth := r.Header.Get("Authorization")
		if auth != fmt.Sprintf("Bearer %s", authKey) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
