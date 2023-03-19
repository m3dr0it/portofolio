package router

import (
	"html/template"
	"net/http"
	"path"
)

func InitViewRoute() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	var errorPagePath = path.Join("view", "error.html")

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var filepath = path.Join("view", "login.html")
			var tmpl, err = template.ParseFiles(filepath)

			if err != nil {
				tmpl.Execute(w, errorPagePath)
				return
			}

			err = tmpl.Execute(w, filepath)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		case "POST":

		}
	})
}
