// Authorization and REgistration handlers
package authorization
import (
	"net/http"
	"web/session"
	"model"
	"logger"
	"errors"
)

type AuthorizeError struct {
	ErrorCode int
	Error     error
}

var CredentialsIncorrectError = errors.New("Username and/or password incorrect.")

// Authorize User
// User password should be plain(not encrypted!) to validate it properly
// It populates the session with a user ID to allow to check future requests against the database.
func AuthorizeUser(w http.ResponseWriter, r *http.Request, user *model.User) *AuthorizeError {
	logger.Info.Print("Trying to Authorize User %+v", user)

	if len(user.Email) == 0 || len(user.Password) == 0 {
		return &AuthorizeError{http.StatusBadRequest, CredentialsIncorrectError}
	}

	dbUser := model.User{
		Email: user.Email,
	}

	if err := dbUser.LoadByEmail(); err != nil {
		return &AuthorizeError{http.StatusBadRequest, err}
	}

	// compare password hashes
	logger.Info.Printf("Authentification User password3. %+v", string(dbUser.Password))
	if err := dbUser.ValidatePassword(user.Password); err != nil {
		return &AuthorizeError{http.StatusBadRequest, CredentialsIncorrectError}
	}

	if err := sessions.SaveUserInSession(w, r, user); err != nil {
		return &AuthorizeError{http.StatusInternalServerError, err}
	}

	return nil
}

// Register new User
func RegisterUser(w http.ResponseWriter, r *http.Request, user *model.User) *AuthorizeError {
	if err := user.Store(); err != nil {
		return &AuthorizeError{http.StatusInternalServerError, err }
	}
	return nil;
}
