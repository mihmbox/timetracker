package controllers

import (
	"github.com/gorilla/mux"
	"logger"
	"middleware"
	"net/http"
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
	ExecuteTemplate(w, "signin", nil)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.AuthentificateUser(w, r)
	if err != nil {
		// Authentication failed, redirect to sign-in page
		logger.Info.Println("Authentification failed. Redirect to ", r.RequestURI)
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
