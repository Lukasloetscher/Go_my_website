package server

import (
	"net/http"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/Lukasloetscher/Go_my_website/pkg/handlers"
	generichandlers "github.com/Lukasloetscher/Go_my_website/pkg/handlers/generic_handlers"
	"github.com/Lukasloetscher/Go_my_website/pkg/mymiddleware"
	"github.com/go-chi/chi"
)

// Create_and_Start_Server() creates a new server and starts listening to the specified Port.
func Create_and_Start_Server(app_ptr *appconfig.AppConfig) error {

	mux := chi.NewRouter()

	mymiddleware.Include_Middleware(app_ptr, mux)

	err := handlers.Add_Manual_Handlers(app_ptr, mux)
	if err != nil {
		panic(err)
	}

	err = generichandlers.Add_Generic_Handlers(app_ptr, mux)
	if err != nil {
		panic(err)
	}
	srv := Create_Server(app_ptr, mux)

	go srv.ListenAndServe()
	return nil
}

// Create_Server() Creates the server
func Create_Server(app_ptr *appconfig.AppConfig, mux *chi.Mux) *http.Server {
	srv := &http.Server{
		Addr:    app_ptr.Port,
		Handler: mux,
	}
	return srv
}
