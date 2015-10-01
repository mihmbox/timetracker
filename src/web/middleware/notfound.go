package middleware

import (
	"app"
	"logger"
	"net/http"
"strings"
)

// Resource not found web.middleware.
// Returns 404 page
func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("NotFound web.middleware", r.URL)
		if !strings.HasPrefix(r.RequestURI, "/api/") {
			// Don't return 404 page if it's API method
			app.App.Template.ExecuteTemplate(w, "404", nil)
		}
	})
}
