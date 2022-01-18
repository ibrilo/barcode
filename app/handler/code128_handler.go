package handler

import (
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	code "github.com/boombuler/barcode/code128"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func Code128(app Application) http.HandlerFunc {
	app.Log().Println("-- Handler Code128 loaded")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		barInt, err := code.Encode(vars["code"])
		if err != nil {
			app.Log().Println(errors.Wrap(err, "create new barcode error"))
			return
		}

		bar, err := barcode.Scale(barInt, 200, 100)
		if err != nil {
			app.Log().Println(errors.Wrap(err, "scale barcode error"))
			return
		}

		png.Encode(w, bar)
		w.Header().Set("Content-Type", "image/png")
	})
}
