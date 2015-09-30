package middleware

import (
	"errors"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"logger"
	"model"
	"net/http"
	"strings"
	"web/session"
)

var CredentialsIncorrectError = errors.New("Username and/or password incorrect.")

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

// AuthenticateUser authenticates a user against the database.
// It populates the session with a user ID to allow web.middleware to check future requests against the database.
func AuthentificateUser(w http.ResponseWriter, r *http.Request) (int, error) {
	logger.Info.Print("Trying to Authentificate User")

	vars := mux.Vars(r)
	email, _ := vars["email"]
	password, _ := vars["password"]
	if len(email) == 0 || len(password) == 0 {
		return http.StatusBadRequest, CredentialsIncorrectError
	}

	logger.Info.Printf("User email: %v, password: %v", email, password)

	user := &model.User{}
	err := user.LoadByEmail(email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Re-direct back to the login page if the user does not exist
	if len(user.Email) == 0 {
		return http.StatusBadRequest, err
	}

	// Leverage the bcrypt package's secure comparison.
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = sessions.SaveUserInSession(w, r, user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
