package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ibrilo/barcode/app/handler"
	"github.com/spf13/viper"
)

type Application struct {
	conf   *viper.Viper
	server *http.Server
	router *mux.Router
	log    *log.Logger
}

func (app *Application) Conf() *viper.Viper   { return app.conf }
func (app *Application) Server() *http.Server { return app.server }
func (app *Application) Router() *mux.Router  { return app.router }
func (app *Application) Log() *log.Logger     { return app.log }

func NewApplication(conf *viper.Viper) *Application {
	return &Application{
		conf: conf,
		log:  log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (app *Application) Run() error {
	app.router = mux.NewRouter()
	addr := fmt.Sprintf("%s:%d",
		app.conf.GetString("server.host"), app.conf.GetInt("server.port"))
	app.server = &http.Server{
		Addr:    addr,
		Handler: app.router,
	}

	app.addHandler("/", handler.Index(app))

	app.log.Printf("Приложение доступно по адресу: http://%s", addr)
	return app.server.ListenAndServe()
}

func (app *Application) addHandler(path string, h http.HandlerFunc) {
	app.router.HandleFunc(path, h)
}
