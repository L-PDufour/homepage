package middleware

import (
	"log"
	"net/http"
)

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Enabling CORS")
		next.ServeHTTP(w, r)
	})
}

func CheckPermissions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Checking Permissions")
		next.ServeHTTP(w, r)
	})
}
