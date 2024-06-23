package routes

import (
	"github.com/adiozdaniel/ascii-art/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlers.HomeHandler(w, r)
		} else if r.FormValue("textInput") != "" {
			handlers.SubmitHandler(w, r)
		} else if r.FormValue("textInput") == "" {
			handlers.BadRequestHandler(w, r)
		} else {
			handlers.NotFoundHandler(w, r)
		}
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/favicon.ico")
	})
}
