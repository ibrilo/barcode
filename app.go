package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type application struct {
	conf   *viper.Viper
	server *http.Server
	router *mux.Router
}

func (app *application) run() error {
	app.router = mux.NewRouter()
	app.server = &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			app.conf.GetString("server.host"),
			app.conf.GetInt("server.port")),
		Handler: app.router,
	}

	return app.server.ListenAndServe()
}
