package app
import (
	"github.com/gorilla/mux"
	"html/template"
	"path/filepath"
	"strings"
	"os"
)

type Application struct {
	Router   *mux.Router
	Template *template.Template
}

var App *Application

func Init() {
	App = &Application{}
	if err := App.LoadTemplates(); err != nil {
		panic("Can not load templates")
	}
}


func (application *Application) LoadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	err := filepath.Walk("./src/views", fn)

	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	application.Template = template.Must(tmpl, err)

	return nil
}