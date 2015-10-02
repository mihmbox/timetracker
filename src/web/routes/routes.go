package routes

import (
	"github.com/gorilla/mux"
	homeCtrl "web/controllers/home"
)

const (
	HOME = "home"
)

func InitRoutes(mx *mux.Router) {
	mx.HandleFunc("/", homeCtrl.Home).Name(HOME)

	mx.HandleFunc("/signin", homeCtrl.SigninPage).Methods("GET")
	mx.HandleFunc("/signin", homeCtrl.Signin).Methods("POST")

	mx.HandleFunc("/signup", homeCtrl.SignupPage).Methods("GET")
	mx.HandleFunc("/signup", homeCtrl.Signup).Methods("POST")
	mx.HandleFunc("/api/signup/validate_email/{email}", homeCtrl.SignupValidateEmail)
}
