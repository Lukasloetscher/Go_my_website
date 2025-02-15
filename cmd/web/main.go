package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// Get_Port_From_Env() gets thed port from the envirement. Used for docker.
// TODO move into a proper package...
func Get_Port_From_Env() string {
	val, ok := os.LookupEnv("PORT")
	if ok {
		return val
	} else {
		return "8080" //default case, if the port is not defined in the envirement.
	}
}

func main() {

	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("usghskjghsehg")) })

	Port := Get_Port_From_Env()

	srv := &http.Server{
		Addr:    ":" + Port,
		Handler: mux,
	}
	os.Getenv("PORT")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
