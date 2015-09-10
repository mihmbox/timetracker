package controllers

import (
	"app"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// Home page Controller
func Home(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"Home page"}

	err := app.App.Template.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SighIn Controller
func Signin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, ok := vars["email"]
	if ok && strings.Index(email, "@") < 0 {
		http.Redirect(w, r, "/", 302)
	}

	w.Write([]byte(email))
}
