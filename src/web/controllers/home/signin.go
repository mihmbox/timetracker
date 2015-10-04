package home

import (
	"logger"
	"net/http"
	"strconv"
	"web/session"
	"web/controllers"
	"web/authorization"
	"model"
)

func SigninPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Email  string
		Failed bool
	}{"", false}

	session, _ := sessions.GetSession(r)
	if errors := session.Flashes(); len(errors) > 0 {
		logger.Info.Println("Flashes count: ", strconv.Itoa(len(errors)))
		session.Save(r, w)
		// there is error flash
		data.Failed = true
		data.Email = r.FormValue("email")
	}

	controllers.ExecuteTemplate(w, "signin", data)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	user := &model.User{
		Email: r.FormValue("email"),
		Password: []byte(r.FormValue("password")),
	}

	err := authorization.AuthorizeUser(w, r, user)
	if err != nil {
		// Authentication failed
		logger.Info.Printf("Authentification failed. Error: %+v", err)

		session, _ := sessions.GetSession(r)
		session.AddFlash(err.Error.Error())
		session.Save(r, w)

		SigninPage(w, r)
		//http.Redirect(w, r, r.Referer(), 302)
	} else {
		targetUrl := r.FormValue("r")
		if len(targetUrl) == 0 {
			targetUrl = "/dashboard"
		}

		http.Redirect(w, r, targetUrl, 302)
	}
}
