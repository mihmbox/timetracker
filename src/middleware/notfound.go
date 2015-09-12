package middleware

import (
	"app"
	"logger"
	"net/http"
)

// Resource not found middleware.
// Returns 404 page
func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("NotFound middleware", r.URL)
		app.App.Template.ExecuteTemplate(w, "404", nil)
		// h.ServeHTTP(w, r)
	})
}
