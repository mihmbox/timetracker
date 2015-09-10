package middleware

import (
	"app"
	"logger"
	"net/http"
)

// Reloads templates to support "live-reload" for debug mode
func ReloadTemplates(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.App.LoadTemplates()
		logger.Info.Println("ReloadTemplates middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}
