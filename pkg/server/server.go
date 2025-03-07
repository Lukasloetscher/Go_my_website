package server

import (
	"log"
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

	go Start_Server(app_ptr, srv)
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
	var err error
	if app_ptr.SecureWebpage {
		log.Println("started secure https server")
		err = srv.ListenAndServeTLS(app_ptr.SecureWebpageCerts.Location_certFile, app_ptr.SecureWebpageCerts.Location_keyFile)

	} else {
		log.Println("started server with http instead of https")
		err = srv.ListenAndServe()
	}
	log.Println(err)
	app_ptr.Channel_Server_Restart <- err
	//TODO if this code is executed, we want to receive an email or something, so we know that we need to restart the server,
	//Also note that such a way to inform us, is a bit problematic, since we can not be sure that, this email arrives in each case...
	//maybe we should instead surveil the container.
}
