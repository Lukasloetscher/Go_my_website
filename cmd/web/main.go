package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

func main() {

	var app_ptr *appconfig.AppConfig
	app_ptr, err := appconfig.Initialise_App_Config()
	if err != nil {
		log.Fatal("error with initialising_App_Config")
	}

	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("called Get function")
		w.Write([]byte("usghskjghsehg"))
	})

	srv := &http.Server{
		Addr:    app_ptr.Port,
		Handler: mux,
	}
	os.Getenv("PORT")
	err = srv.ListenAndServe()
	log.Fatal(err)
}
