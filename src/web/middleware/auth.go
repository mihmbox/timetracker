package middleware

import (
	"logger"
	"net/http"
	"strings"
	"web/session"
)

// Authorization web.middleware.
// Redirects to sign-in screen if user is not authorized.
func AuthMiddleware(h http.Handler, prefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip authorization
		if !strings.HasPrefix(r.RequestURI, prefix) {
			h.ServeHTTP(w, r)
			return
		}

		logger.Info.Println("Authorization web.middleware", r.RequestURI)
		user, err := sessions.GetUserFromSession(r)
		if err != nil {
			// can't get Session. Log and redirect to Sign-in page
			logger.Error.Print("Cannot get user from session: " + err.Error() + user.Email)
		}

		if err != nil || len(user.Email) == 0 {
			// User is not logged in -> redirect to Sign-In page
			http.Redirect(w, r, "/signin?r="+r.RequestURI, 302)
			return
		}

		logger.Info.Println("User id= ", user.ID)

		// User is authenticated
		h.ServeHTTP(w, r)
	})
}