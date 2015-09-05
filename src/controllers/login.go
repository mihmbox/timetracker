package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	//"routes"
	"strings"
)

const (
	HOME = "home"
	SIGNIN = "signin"
)


func Build(mx *mux.Router)  {
	mx.HandleFunc("/", Home).Name(HOME)
	mx.HandleFunc("/signin/{email}", Signin).Name(SIGNIN) //.Methods("POST")
}


func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"));
}

func Signin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, ok := vars["email"]
	if ok && strings.Index(email, "@") < 0 {
		http.Redirect(w, r, "/", 302)
	}

	w.Write([]byte(email));
}