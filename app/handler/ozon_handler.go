package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func Ozon(app Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl, err := loadTemplate(filepath.Join(app.Conf().GetString("templatesDir"), "ozon_barcode.html"))
		if err != nil {
			app.Log().Println(errors.Wrap(err, "tempalte parse error"))
			w.Write([]byte("template parse eror"))
			return
		}

		vars := mux.Vars(r)

		data := struct {
			Code string
		}{vars["code"]}

		if err := tpl.Execute(w, data); err != nil {
			app.Log().Println(errors.Wrap(err, "template execute error"))
			w.Write([]byte("template execute error"))
			return
		}
	})
}
