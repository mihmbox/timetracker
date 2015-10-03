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
	"strings"
	"regexp"
)

var (
	EmailExistsOrInvalidErrorCode = 1
	PasswordIsWeakErrorCode = 2
	PasswordsDontMatchErrorCode = 3
	RegistrationInternalServerError = 4
)
// Shows Signup page.
func SignupPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		User      model.User
		ErrorCode int
	}{model.User{}, 0}

	if userObj, ok := context.GetOk(r, "user"); ok {
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

// Validate user input, register new user and redirects to Dashboard
func Signup(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	password := strings.TrimSpace(r.FormValue("password"))
	passwordConfirm := r.FormValue("passwordconfirm")
	user := model.User{
		Email: strings.TrimSpace(r.FormValue("email")),
		Password: []byte(password),
	}
	// Save User in context to use it on Signup page later if registration failed
	context.Set(r, "user", user)

	// Registration failed handler
	failed := func(errCode int) {
		logger.Info.Printf("Registration failed.Error: %+v, user: %+v", errCode, user)
		session, _ := sessions.GetSession(r)
		session.AddFlash(errCode)
		session.Save(r, w)

		SignupPage(w, r)
	}

	// Validation: email
	emailIsValid := validateEmail(user.Email) == http.StatusOK;
	if !emailIsValid {
		failed(EmailExistsOrInvalidErrorCode)
		return
	}

	// Validation: if Password is weak
	passwordIsWeak := len(password) <= 6
	if passwordIsWeak {
		failed(PasswordIsWeakErrorCode)
		return
	}

	// Validation: compare passwords
	passwordsDifferent := password != passwordConfirm
	if passwordsDifferent {
		failed(PasswordsDontMatchErrorCode)
		return
	}

	if err := authorization.RegisterUser(w, r, &user); err != nil {
		failed(RegistrationInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", 302)
}

// Handler for "/api/signup/validate_email/{email}"
func SignupValidateEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	respStatus := validateEmail(email)
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

// Validate Email. Returns '200' http status Ok if email is valid and available
func validateEmail(email string) int{
	var respStatus int
	// check if string is email
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(email))
	if matched == false {
		respStatus = http.StatusBadRequest
	}

	// check if email exists in database
	user := model.User{}
	if err := user.LoadByEmail(user.Email); err == nil {
		// user already exists with specified email
		respStatus = http.StatusForbidden
	}

	if respStatus <= 0 {
		// set "Ok" status only if all validation steps passed
		respStatus = http.StatusOK
	}

	return respStatus
}