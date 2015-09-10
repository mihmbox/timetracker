package main

import (
	"app"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"logger"
	"middleware"
	"model"
	"net/http"
	"os"
	"routes"
)

func init() {
	// Configure App logger
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	flag.Parse()
}

func main() {
	app.Init()

	// Build DB Model
	if app.App.Config.Db.CreateDB {
		logger.Info.Println("Starting DB creation")
		if err := model.CreateDB(); err != nil {
			logger.Error.Println("Cannot create DB")
			panic(err.Error())
		} else {
			logger.Info.Println("DB created")
		}
	}

	// Build routes and start webApp
	router := mux.NewRouter()
	// Routes
	routes.Init(router)
	// Static files
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	var handler http.Handler = router
	// Logging
	handler = handlers.LoggingHandler(os.Stdout, handler)
	// Authorization middleware
	handler = middleware.AuthMiddleware(handler)
	// liveReload
	if app.App.Config.Env.DevMode {
		handler = middleware.ReloadTemplates(handler)
	}

	port := app.App.Config.Server.Port
	fmt.Println("Listening. Port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), handler)
}
