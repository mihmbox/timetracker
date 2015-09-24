package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
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
	vars := mux.Vars(r)
	email, ok := vars["email"]
	if ok && strings.Index(email, "@") < 0 {
		http.Redirect(w, r, "/", 302)
	}

	w.Write([]byte(email))
}
