package middleware

import (
	"app"
	"logger"
	"net/http"
)

// Reloads templates to support "live-reload" for debug mode
func ReloadTemplates(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("ReloadTemplates web.middleware", r.URL)
		if err := app.App.LoadTemplates(); err != nil {
			logger.Error.Println("Cannot reload tempalates", err.Error())
		}

		h.ServeHTTP(w, r)
	})
}
