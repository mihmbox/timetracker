package app

import (
	"model"
	"net/http"
)

func GetUserFromSession(r *http.Request) (*model.User, error) {
	session, err := App.GetSession(r)
	if err != nil {
		return nil, err
	}

	userId := session.Values["userId"]
	email := session.Values["email"]
	network := session.Values["network"]

	return &model.User{
		ID:      userId.(int),
		Email:   email.(string),
		Network: network.(string),
	}, nil

}

func SaveUserInSession(w http.ResponseWriter, r *http.Request, user *model.User) error {
	session, err := App.GetSession(r)
	if err != nil {
		return nil, err
	}

	session.Values["userId"] = user.ID
	session.Values["email"] = user.Email
	session.Values["network"] = user.Network

	err = session.Save(r, w)
	return err
}
