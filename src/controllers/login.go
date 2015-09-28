package controllers

import (
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"middleware"
	"model"
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
	_, err := middleware.AuthenticateUser(w, r)
	if err != nil {
		// authentification failed, redirect to sign-in page
		http.Redirect(w, r, "/signin", http.StatusUnauthorized)
	} else {
		vars := mux.Vars(r)
		targetUrl := vars["p"]
		if len(targetUrl) == 0 {
			targetUrl = "/dashboard"
		}

		http.Redirect(w, r, targetUrl, 302)
	}
}
