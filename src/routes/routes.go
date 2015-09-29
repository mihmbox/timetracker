package routes

import (
	"controllers"
	"github.com/gorilla/mux"
)

const (
	HOME = "home"
)

func InitRoutes(mx *mux.Router) {
	mx.HandleFunc("/", controllers.Home).Name(HOME)

	mx.HandleFunc("/signin", controllers.SigninPage).Methods("GET")
	mx.HandleFunc("/signin", controllers.Signin).Methods("POST")

	mx.HandleFunc("/signup", controllers.SignupPage)
}
