package controllers

import (
	"app"
	"net/http"
)

func ExecuteTemplate(wr http.ResponseWriter, template string, data interface{}) {
	err := app.App.Template.ExecuteTemplate(wr, template, data)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}
