package controllers

import (
	"app"
	"net/http"
	"logger"
)

func ExecuteTemplate(wr http.ResponseWriter, template string, data interface{}) {

	// Reload template to have latest changes only in Dev mode
	if app.App.Config.Env.DevMode {
		if err := app.App.LoadTemplates(); err != nil {
			logger.Error.Println("Cannot reload tempalates", err.Error())
		}
	}

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