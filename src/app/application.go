package app

import (
	"code.google.com/p/gcfg"
	"code.google.com/p/gorilla/sessions"
	"github.com/gorilla/mux"
	"html/template"
	"logger"
	"os"
	"path/filepath"
	"strings"
)

type Application struct {
	Config       Config
	Router       *mux.Router
	Template     *template.Template
	SessionStore *sessions.CookieStore
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

	App.SessionStore = sessions.NewCookieStore([]byte(App.Config.Server.SessionKey))
	//	App.SessionStore.Options = &session.Options{
	//		Path:     "/",
	//		MaxAge:   60*30, // 30 minutes
	//		Secure:   true,
	//	}
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

	err := filepath.Walk(app.Config.Server.Templates, fn)

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
