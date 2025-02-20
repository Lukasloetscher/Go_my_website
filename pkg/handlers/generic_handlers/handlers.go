package generichandlers

import (
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/Lukasloetscher/Go_my_website/pkg/models"
	"github.com/Lukasloetscher/Go_my_website/pkg/render"
	"github.com/go-chi/chi"
)

// Add_Generic_Handlers()
// adds the handlers to the code, whicha re generated automatically, based on the files in the specifided fodler
func Add_Generic_Handlers(app_ptr *appconfig.AppConfig, mux *chi.Mux) error {

	err := filepath.WalkDir(app_ptr.Webpage_Location, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".generic_page.tmpl") {
			name, _ := strings.CutSuffix(path, ".generic_page.tmpl")
			name = strings.ReplaceAll(name, "\\", "/")
			name, _ = strings.CutPrefix(name, app_ptr.Webpage_Location)

			mux.Get(name, func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					r := recover()
					if r != nil {
						log.Println("Recovcered from Panic")
						log.Println(r)
						w.Write([]byte("405 Internal error"))
					}
				}()
				//Here we need to render the page...
				var paths render.Render_Filepaths
				paths.Filepath_Html = path
				paths.Filepath_Layout = app_ptr.Layout_Location

				var data models.TemplateData
				data.BoolMap = make(map[string]bool)
				data.BoolMap["navbar"] = true
				buf := render.Create_Rendered_Template(paths, &data)
				buf.WriteTo(w)
			})

		}
		return nil
	})

	return err
}
