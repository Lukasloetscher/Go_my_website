package main

import (
	"log"
	"net/http"
	"os"

	getport "github.com/Lukasloetscher/Go_my_website/util/get_port"
	"github.com/go-chi/chi"
)

func main() {

	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("called Get function")
		w.Write([]byte("usghskjghsehg"))
	})

	Port := getport.Get_Port_From_Env() //Todo we later need to handle panics!

	srv := &http.Server{
		Addr:    ":" + Port,
		Handler: mux,
	}
	os.Getenv("PORT")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
