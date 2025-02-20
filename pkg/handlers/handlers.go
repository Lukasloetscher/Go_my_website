package handlers

import (
	"log"
	"net/http"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/Lukasloetscher/Go_my_website/pkg/models"
	"github.com/Lukasloetscher/Go_my_website/pkg/render"
	"github.com/go-chi/chi"
)

// Add_Manual_Handlers()
// adds the handlers to the code, which i write manually
func Add_Manual_Handlers(app_ptr *appconfig.AppConfig, mux *chi.Mux) error {

	//for testing only. ove to better place!
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Println("Recovcered from Panic")
				log.Println(r)
				w.Write([]byte("405 Internal error"))
			}
		}()
		var paths render.Render_Filepaths
		paths.Filepath_Html = "webpages\\main.tmpl"
		paths.Filepath_Layout = ""
		var data *models.TemplateData
		buf := render.Create_Rendered_Template(paths, data)
		buf.WriteTo(w)
	})

	return nil
}
