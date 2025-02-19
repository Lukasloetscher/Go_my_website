package handlers

import (
	"net/http"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

// Add_Manual_Handlers()
// adds the handlers to the code, which i write manually
func Add_Manual_Handlers(app_ptr *appconfig.AppConfig, mux *chi.Mux) error {

	//For the moment, i just add a simple palcehodler, so we can still see something:
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//Here we need to render the page...
		w.Write([]byte("Hallo Welt"))
	})

	return nil
}
