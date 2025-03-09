package getport

import (
	"os"
	"regexp"
)

const regex_port string = "^[0-9]+$"

// check_If_Valid_Port checks if the entry is a valid port umber, i.e. it needs to be numeric.
// Currently we do not test, wheter it is a reasonable port number
func check_If_Valid_Port(port string) (ok bool) {
	matched, err := regexp.MatchString(regex_port, port)
	if err != nil {
		panic(err)
	} else {
		return matched

	}
}

const port_env string = "PORT"

// Get_Port_From_Env() gets the port from the environment. Used for docker.
// When the env variable is not set, it defaults to 8080
// due tio the way ports are handled in docker, it does not make sense to use this. we need to expose the prot anyway in the dockerfile...
func Get_Port_From_Env() (port string) {
	port, ok := os.LookupEnv(port_env)
	if ok {
		if check_If_Valid_Port(port) {
			return port
		} else {
			panic("No Valid Portnumber was entered.")
		}
	} else {
		return "8080" //default case, if the port is not defined in the environment.
	}
}
