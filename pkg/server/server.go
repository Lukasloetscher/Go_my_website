package server

import (
	"net/http"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

// Create_and_Start_Server() creates a new server and starts listening to the specified Port.
func Create_and_Start_Server(app_ptr *appconfig.AppConfig) error {

	mux := chi.NewRouter()
	//Add_Middleware to server

	//Add manual Handlers to Server.

	//Add generic Routes

	//Create Server
	srv := Create_Server(app_ptr, mux)

	//start to listen to server.
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
