package home

import (
	"net/http"
	"web/controllers"
	"github.com/gorilla/mux"
	"logger"
	"web/session"
	"web/authorization"
	"model"
	"github.com/gorilla/context"
)

var (
	EmailExistsOrInvalidErrorCode = 1
	PasswordsDontMatchErrorCode = 2
)
// Shows Signup oag
func SignupPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		User      model.User
		ErrorCode int
	}{model.User{}, 0}

	if userObj, ok := context.GetOk(r, "user"); ok{
		data.User, _ = userObj.(model.User)
	}

	session, _ := sessions.GetSession(r)
	if errors := session.Flashes(); len(errors) > 0 {
		session.Save(r, w)
		data.ErrorCode = EmailExistsOrInvalidErrorCode
		if errCode, ok := errors[0].(int); ok {
			data.ErrorCode = errCode
		}
	}

	controllers.ExecuteTemplate(w, "signup", data)
}

// Executes SignUp logic
func Signup(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	user := model.User{
		Email: r.FormValue("email"),
		Password: []byte(r.FormValue("password")),
	}
	//passwordConfirm := r.FormValue("passwordconfirm")

	//	re := regexp.MustCompile(".+@.+\\..+")
	//	matched := re.Match([]byte(msg.Email))
	//	if matched == false {
	//		msg.Errors["Email"] = "Please enter a valid email address"
	//	}
	context.Set(r, "user", user)
	if err := authorization.RegisterUser(w, r, &user); err != nil {
		// Registration failed
		logger.Info.Printf("Registration failed.Error: %+v, user: %+v", err, user)
		session, _ := sessions.GetSession(r)
		session.AddFlash(err.Error.Error())
		session.Save(r, w)

		SignupPage(w, r)
	} else {
		http.Redirect(w, r, "/dashboard", 302)
	}
}

// Handler for "/api/signup/validate_email/{email}"
func SignupValidateEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	isValid := email == "mik@test.com"

	respStatus := http.StatusForbidden
	if isValid {
		respStatus = http.StatusOK
	}

	w.WriteHeader(respStatus)

	//	responseJson := struct {
	//		Success bool
	//	}{isValid}
	//	jsonStr, err := json.Marshal(responseJson)
	//	if err != nil {
	//		controllers.RespondError(w, err)
	//		return
	//	}
	//
	//	controllers.RespondJSON(w, jsonStr)
}