package controllers

import (
	"app"
	"net/http"
)

func ExecuteTemplate(wr http.ResponseWriter, template string, data interface{}) {
	err := app.App.Template.ExecuteTemplate(wr, template, data)
	if err != nil {
		RespondError(wr,err)
	}
}

// Sends InternalServer error response
func RespondError(wr http.ResponseWriter, err error) {
	http.Error(wr, err.Error(), http.StatusInternalServerError)
}

// Writes JSON in Response
func RespondJSON(wr http.ResponseWriter, json []byte) {
	wr.Header().Set("Content-Type", "application/json; charset=utf-8")
	wr.Write(json)
}