package middleware

import (
	"context"
	"homepage/internal/auth"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}
type contextKey string

const userContextKey contextKey = "user"

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func WithAuthenticator(authenticator *auth.Authenticator) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var user *auth.AuthenticatedUser
			var err error

			// Attempt to verify the token based on the environment
			if auth.IsProduction {
				// Production: verify the token
				user, err = authenticator.VerifyToken(r)
			} else {
				// Non-production: use mock verification
				user, err = authenticator.MockVerifyToken(r)
			}

			if err != nil {
				// Log the failure and proceed as unauthenticated
				log.Printf("Failed to verify token: %v. Proceeding as unauthenticated.\n", err)
			} else {
				// Token is valid, add the authenticated user to the context
				log.Printf("Token verified successfully for user: %+v\n", user)
				ctx := context.WithValue(r.Context(), userContextKey, user)
				r = r.WithContext(ctx)
			}

			// Proceed with or without the user in the context
			next.ServeHTTP(w, r)
		})
	}
}

func GetUserFromContext(ctx context.Context) (*auth.AuthenticatedUser, bool) {
	user, ok := ctx.Value(userContextKey).(*auth.AuthenticatedUser)
	return user, ok
}
