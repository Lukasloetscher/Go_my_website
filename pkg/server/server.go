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

	/*
	   TODO move to correct place (create a function)
	   TODO find out where to save images. I'm not sure, we should keep them in the docker file...s
	*/
	fileServer := http.FileServer(http.Dir("./webpages/0_static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

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

// Start_Server() Starts the Server. Depending on teh appconfig setting either in http or https mode.
func Start_Server(app_ptr *appconfig.AppConfig, srv *http.Server) {
	if app_ptr.SecureWebpage {
		go srv.ListenAndServeTLS("todo", "todo")
	} else {
		go srv.ListenAndServe()
	}
}
