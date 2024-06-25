package routes

import (
	"net/http"

	"github.com/adiozdaniel/ascii-art/handlers"
	"github.com/adiozdaniel/ascii-art/utils"
)

//RegisterRoutes manages the routes
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.Inputs.Input = r.FormValue("textInput")
		banner := utils.BannerFiles[r.FormValue("FileName")]
		if banner == "" {
			utils.Inputs.BannerPath = utils.BannerFiles["standard"]
		} else {
			utils.Inputs.BannerPath = utils.BannerFiles[r.FormValue("FileName")]
		}

		if r.URL.Path == "/" {
			handlers.HomeHandler(w, r)
		} else if utils.Inputs.Input != "" && r.FormValue("FileName") != "" {
			handlers.SubmitHandler(w, r)
		} else if utils.Inputs.Input == "" && r.URL.Path == "/ascii-art" {
			handlers.BadRequestHandler(w, r)
		} else if utils.Inputs.Input != "" && r.FormValue("FileName") == "" {
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
