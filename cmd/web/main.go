package main

import (
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/go-chi/chi"
)

const regex_port string = "^[0-9]+$"

// Check_If_Valid_Port checks if the entry is a valid port umber, i.e. it needs to be numeric.
// Currently we do not test, wheter it is a reasonable port number
// TODO move into a proper package...
// When moving this should include the test function
func Check_If_Valid_Port(port string) (ok bool) {
	matched, err := regexp.MatchString(regex_port, port)
	if err != nil {
		panic(err)
	} else {
		return matched

	}
}

const port_env string = "PORT"

// Get_Port_From_Env() gets thed port from the envirement. Used for docker.
// After reading up on how port forwarding works in docker, this code is a bit redundant.
// That beeing said, i still leave it here, so i can check if the LookupEnv works in docker as i expect.
// Also there is no real harm in it.
// I might delete this later.
// TODO move into a proper package...
func Get_Port_From_Env() (port string) {
	port, ok := os.LookupEnv(port_env)
	if ok {
		if Check_If_Valid_Port(port) {
			return port
		} else {
			panic("No Valid Portnumber was entered.")
		}
	} else {
		return "8080" //default case, if the port is not defined in the envirement.
	}
}

func main() {

	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("called Get function")
		w.Write([]byte("usghskjghsehg"))
	})

	Port := Get_Port_From_Env() //Todo we later need to handle panics!

	srv := &http.Server{
		Addr:    ":" + Port,
		Handler: mux,
	}
	os.Getenv("PORT")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
