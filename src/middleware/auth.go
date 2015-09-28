package middleware

import (
	"app"
	"errors"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"logger"
	"model"
	"net/http"
)

var CredentialsIncorrectError = errors.New("Username and/or password incorrect.")

// Authorization middleware.
// Redirects to sign-in screen if user is not authorized.
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Authorization middleware", r.URL)

		user, err := app.GetUserFromSession(r)
		if err != nil {
			// can't get Session. Log and redirect to Sign-in page
			logger.Error.Print("Cannot get user from session: " + err)
			http.Redirect(w, r, "/signin", 302)
			return
		}
		if len(user.Email) == 0 {
			// User is not logged in -> redirect to Sign-In page
			http.Redirect(w, r, "/signin", 302)
			return
		}

		// User is authenticated
		h.ServeHTTP(w, r)
	})
}

// AuthenticateUser authenticates a user against the database.
// It populates the session with a user ID to allow middleware to check future requests against the database.
func AuthenticateUser(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	email, _ := vars["email"]
	password, _ := vars["password"]
	if len(email) == 0 || len(password) == 0 {
		return http.StatusBadRequest, CredentialsIncorrectError
	}

	user := &model.User{}
	err := user.LoadByEmail(email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Re-direct back to the login page if the user does not exist
	if len(user.Email) == 0 {
		//		// Save error in session flash
		//		session.AddFlash(ErrCredentialsIncorrect, "_errors")
		//		err := session.Save(r, w)
		//		if err != nil {
		//			return 500, err
		//		}
		//		http.Redirect(w, r, loginURL, 302)
		//		return 302, err
		return http.StatusBadRequest, err
	}

	// Leverage the bcrypt package's secure comparison.
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = app.SaveUserInSession(w, r, user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
