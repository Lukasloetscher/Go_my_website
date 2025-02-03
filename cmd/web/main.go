package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("usghskjghsehg")) })

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
