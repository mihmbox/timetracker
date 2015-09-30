package sessions

import (
	"app"
	"code.google.com/p/gorilla/sessions"
	"logger"
	"model"
	"net/http"
)

const userSessionKey = "timetracker-user"

// Get Session
func GetSession(r *http.Request) (*sessions.Session, error) {
	return app.App.SessionStore.Get(r, "timetracker-user")
}

func GetUserFromSession(r *http.Request) (*model.User, error) {
	session, err := GetSession(r)
	if err != nil {
		return nil, err
	}

	user := &model.User{}

	logger.Info.Println("GetUserFromSession")

	// Parse Session data
	if userId, ok := session.Values["userId"].(int); ok {
		user.ID = userId
	}
	if email, ok := session.Values["email"].(string); ok {
		user.Email = email
	}
	if network, ok := session.Values["network"].(string); ok {
		user.Network = network
	}

	return user, nil
}

func SaveUserInSession(w http.ResponseWriter, r *http.Request, user *model.User) error {
	session, err := GetSession(r)
	if err != nil {
		return err
	}

	session.Values["userId"] = user.ID
	session.Values["email"] = user.Email
	session.Values["network"] = user.Network

	err = session.Save(r, w)
	return err
}
