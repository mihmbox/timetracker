package home

import (
	"net/http"
	"web/controllers"
)

// Home page Controller
func Home(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"Home page"}
	controllers.ExecuteTemplate(w, "index", data)
}
