package main

import (
	"app"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"logger"
	"model"
	"net/http"
	"os"
	"web/routes"
	"web/middleware"
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

	// Build web.routes and start webApp
	router := mux.NewRouter()
	// Routes
	routes.InitRoutes(router)
	// Static files
	router.PathPrefix("/public/").Handler(getStaticFilesHandler())
	router.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules/"))))
	// 404 handler
	router.NotFoundHandler = middleware.NotFound()

	var handler http.Handler = router
	// Authorization web.middleware
	handler = middleware.AuthMiddleware(handler, "/dashboard")
	// Logging
	handler = handlers.LoggingHandler(os.Stdout, handler)
	// Panic recovery
	handler = middleware.PanicRecovery(handler)

	// liveReload for Templates. Put in ExecuteTemplate method
//	if app.App.Config.Env.DevMode {
//		handler = middleware.ReloadTemplates(handler)
//	}

	port := app.App.Config.Server.Port
	fmt.Println("Listening. Port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), handler)
}

// Builds static files handler:
//	- Proxy to webpack server for Dev mode
//	- FileServer for Prod mode
func getStaticFilesHandler() http.Handler {
	isDevMode := app.App.Config.Env.DevMode
	assetsProxyUrl := app.App.Config.Server.AssetsProxy
	if isDevMode && assetsProxyUrl != "" {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "http://localhost:8080"+r.RequestURI, 302)
		})
	} else {
		return http.StripPrefix("/public/", http.FileServer(http.Dir("public/")))
	}
}
