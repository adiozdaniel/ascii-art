package routes

import (
	"net/http"

	"github.com/adiozdaniel/ascii-art/handlers"
)

// RegisterRoutes manages the routes
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})

	mux.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.SubmitHandler(w, r)
		} else {
			handlers.HomeHandler(w, r)
		}
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/favicon.ico")
	})
}
