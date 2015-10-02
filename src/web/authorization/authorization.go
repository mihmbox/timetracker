// Authorization and REgistration handlers
package authorization
import (
	"net/http"
	"web/session"
	"golang.org/x/crypto/bcrypt"
	"model"
	"logger"
	"errors"
)

type AuthorizeError struct {
	ErrorCode int
	Error     error
}

var CredentialsIncorrectError = errors.New("Username and/or password incorrect.")

// AuthorizeUser authorizes user.
// It populates the session with a user ID to allow to check future requests against the database.
func AuthorizeUser(w http.ResponseWriter, r *http.Request, user *model.User) *AuthorizeError {
	logger.Info.Print("Trying to Authentificate User")

	if len(user.Email) == 0 || len(user.Password) == 0 {
		return &AuthorizeError{http.StatusBadRequest, CredentialsIncorrectError}
	}

	logger.Info.Printf("User email: %v, password: %v", user.Email, user.Password)

	err := user.LoadByEmail(user.Email)
	if err != nil {
		return &AuthorizeError{http.StatusInternalServerError, err}
	}

	// Re-direct back to the login page if the user does not exist
	if len(user.Email) == 0 {
		return &AuthorizeError{http.StatusBadRequest, err}
	}

	// Leverage the bcrypt package's secure comparison.
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(user.Password))
	if err != nil {
		return &AuthorizeError{http.StatusBadRequest, err}
	}

	err = sessions.SaveUserInSession(w, r, user)
	if err != nil {
		return &AuthorizeError{http.StatusInternalServerError, err}
	}

	return nil
}

// Register new User
func RegisterUser(w http.ResponseWriter, r *http.Request, user *model.User) *AuthorizeError {
	return &AuthorizeError{http.StatusInternalServerError, errors.New("No implemented")}
}
