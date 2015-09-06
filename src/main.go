package main
import (
	"model"
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"logger"
	"io/ioutil"
	"middleware"
	"flag"
	"routes"
	"app"
)

var (
	port = flag.Int("port,p", 8080, "Http port")
	initDB = flag.Bool("initDB", true, "Crete DB schema")
)


func init() {
	app.Init()

	// Configure App logger
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	flag.Parse()
}


func main() {
	// Build DB Model
	if *initDB {
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

	fmt.Println("Listening. Port: ", *port)
	http.ListenAndServe(fmt.Sprintf(":%v", *port), handler)
}