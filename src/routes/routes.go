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

	mx.HandleFunc("/signin", controllers.SigninPage)
	mx.HandleFunc("/signup", controllers.SignupPage)
	mx.HandleFunc("/signin/{email}", controllers.Signin)
}
