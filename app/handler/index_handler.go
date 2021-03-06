package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Application interface {
	Conf() *viper.Viper
	Server() *http.Server
	Router() *mux.Router
	Log() *log.Logger
}

func Index(app Application) http.HandlerFunc {
	app.Log().Println("-- Handler Index loaded")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
