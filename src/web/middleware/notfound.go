package middleware

import (
	"app"
	"logger"
	"net/http"
)

// Resource not found web.middleware.
// Returns 404 page
func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("NotFound web.middleware", r.URL)
		app.App.Template.ExecuteTemplate(w, "404", nil)
		// h.ServeHTTP(w, r)
	})
}
