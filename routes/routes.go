package routes

import (
	"net/http"

	"github.com/adiozdaniel/ascii-art/handlers"
	"github.com/adiozdaniel/ascii-art/utils"
)

// RegisterRoutes manages the routes
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			utils.Inputs.Input = r.FormValue("textInput")
			banner := utils.BannerFiles[r.FormValue("FileName")]
			if banner == "" {
				utils.Inputs.BannerPath = utils.BannerFiles["standard"]
			} else {
				utils.Inputs.BannerPath = banner
			}

			if utils.Inputs.Input != "" {
				handlers.SubmitHandler(w, r)
			} else {
				handlers.BadRequestHandler(w, r)
			}
		} else {
			handlers.HomeHandler(w, r)
		}
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/favicon.ico")
	})
}
