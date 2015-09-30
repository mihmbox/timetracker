package middleware

import (
	"fmt"
	"logger"
	"net/http"
)

// Panic recovery
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error.Printf("Panic: %+v", err)
				http.Error(w, fmt.Sprintf("Sorry, but server has some problems: \n%+v", err), 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
