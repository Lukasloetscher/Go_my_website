package generichandlers

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/Lukasloetscher/Go_my_website/pkg/models"
	"github.com/Lukasloetscher/Go_my_website/pkg/render"
	htmlerrors "github.com/Lukasloetscher/Go_my_website/util/html_errors"
	"github.com/go-chi/chi"
)

// Add_Generic_Handlers()
// adds the handlers to the code, whicha re generated automatically, based on the files in the specifided fodler
func Add_Generic_Handlers(app_ptr *appconfig.AppConfig, mux *chi.Mux) error {

	err := filepath.WalkDir(app_ptr.Gen_Pages.Webpage_Location, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".generic_page.tmpl") {
			var paths render.Render_Filepaths
			paths.Filepath_Html = path
			paths.Filepath_Layout = app_ptr.Gen_Pages.Layout_Location
			var data models.TemplateData
			data.BoolMap = make(map[string]bool)
			data.BoolMap["navbar"] = app_ptr.Gen_Pages.Add_Navbar

			mux.Get(generate_generic_page_name(path, app_ptr), func(w http.ResponseWriter, r *http.Request) {
				defer htmlerrors.Catch_Panic_And_Report_Internal_Error(w)
				buf := render.Create_Rendered_Template(paths, &data)
				buf.WriteTo(w)
			})

		}
		return nil
	})

	return err
}

func generate_generic_page_name(path string, app_ptr *appconfig.AppConfig) string {
	name, _ := strings.CutSuffix(path, ".generic_page.tmpl")
	name = strings.ReplaceAll(name, "\\", "/")
	name, _ = strings.CutPrefix(name, app_ptr.Gen_Pages.Webpage_Location)
	return name
}
