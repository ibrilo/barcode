package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Application struct {
	conf   *viper.Viper
	server *http.Server
	router *mux.Router
}

func NewApplication(conf *viper.Viper) *Application {
	return &Application{
		conf: conf,
	}
}

func (app *Application) Run() error {
	app.router = mux.NewRouter()
	app.server = &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			app.conf.GetString("server.host"),
			app.conf.GetInt("server.port")),
		Handler: app.router,
	}

	return app.server.ListenAndServe()
}
