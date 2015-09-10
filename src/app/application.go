package app

import (
	"code.google.com/p/gcfg"
	"github.com/gorilla/mux"
	"html/template"
	"logger"
	"os"
	"path/filepath"
	"strings"
)

type Application struct {
	Config   Config
	Router   *mux.Router
	Template *template.Template
}

var App *Application

func Init() {
	App = &Application{}

	if err := App.InitCfg(); err != nil {
		logger.Error.Println(err.Error())
		panic("Can't parse configuration in file './src/config.ini'")
	}

	if err := App.LoadTemplates(); err != nil {
		logger.Error.Println(err.Error())
		panic("Can not load templates")
	}
}

// Parse and load configurations from file
func (app *Application) InitCfg() error {
	var cfg Config
	if err := gcfg.ReadFileInto(&cfg, "./src/config.ini"); err != nil {
		return err
	}

	app.Config = cfg
	return nil
}

// Loads and parses html templates
func (app *Application) LoadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	//	pwd, _ := os.Getwd()
	//	path := pwd + "/views/"
	//	err := filepath.Walk(path, fn)
	err := filepath.Walk(app.Config.Server.Public, fn)

	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	app.Template = template.Must(tmpl, err)

	return nil
}
