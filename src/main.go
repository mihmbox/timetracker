package main
import (
	"model"
	"fmt"
	"controllers"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	CREATE_DB = true
	PORT = 8080
)

func main() {
	// Build DB Model
	if CREATE_DB {
		if err := model.CreateDB(); err != nil {
			panic(err.Error())
		} else {
			fmt.Println("DB created")
		}
	}

	// Build routes and start webApp
	mx := mux.NewRouter()
	controllers.Build(mx)

	// Static files
	mx.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))


	fmt.Println("Listening. Port: ", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), mx)

//	n := negroni.Classic()
//	n.UseHandler(mx)
//	n.Run(fmt.Sprintf(":%v", PORT))
}