package controllers

import (
	"github.com/gorilla/mux"
	"logger"
	"net/http"
	"strconv"
	"web/middleware"
	"web/session"
)

// Home page Controller
func Home(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"Home page"}
	ExecuteTemplate(w, "index", data)
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, "signup", nil)
}
func SigninPage(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("SigninPage ")
	data := struct {
		Error  string
		Failed bool
	}{"", false}

	session, _ := sessions.GetSession(r)
	if errors := session.Flashes(); len(errors) > 0 {
		logger.Info.Println("Flashes count: ", strconv.Itoa(len(errors)))
		session.Save(r, w)
		// there is error flash
		data.Failed = true
		data.Error = errors[0].(string)
	}

	ExecuteTemplate(w, "signin", data)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.AuthentificateUser(w, r)
	if err != nil {
		// Authentication failed, redirect to sign-in page
		logger.Info.Println("Authentification failed. Redirect to ", r.RequestURI)

		session, _ := sessions.GetSession(r)
		session.AddFlash(err.Error())
		session.Save(r, w)

		http.Redirect(w, r, r.Referer(), 302)
	} else {
		vars := mux.Vars(r)
		targetUrl := vars["r"]
		if len(targetUrl) == 0 {
			targetUrl = "/dashboard"
		}

		http.Redirect(w, r, targetUrl, 302)
	}
}
