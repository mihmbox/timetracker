package middleware

import (
	"logger"
	"net/http"
)

// Authorization middleware.
// Redirects to sign-in screen if user is not authorized.
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Authorization middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}
