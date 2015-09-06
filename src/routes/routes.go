package routes

import (
	"github.com/gorilla/mux"
	"controllers"
)

const (
	HOME = "home"
	SIGNIN = "signin"
)

func Init(mx *mux.Router) {
	mx.HandleFunc("/", controllers.Home).Name(HOME)
	mx.HandleFunc("/signin/{email}", controllers.Signin).Name(SIGNIN) //.Methods("POST")
}