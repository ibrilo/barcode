package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func Ozon(app Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templateText := `<!doctype html>
		<html>
			<head>
				<style>
				
				@page {
					size: 58mm 40mm;
				}
  
				// .content {
				// 	width: 58mm;
				// 	height: 40mm;
				// 	display: flex;
				// 	justify-content: center;
				// }
				</style>
			</head>
			<body>
				<div class="content">
					<div>
						<img src="/code128/{{.Code}}">
						<h4>{{.Code}}</h4>
					</div>
				</div>
			</body>
		</html>`

		tpl, err := template.New("page").Parse(templateText)
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
