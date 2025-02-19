package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/Lukasloetscher/Go_my_website/pkg/server"
)

func main() {

	var app_ptr *appconfig.AppConfig
	app_ptr, err := appconfig.Initialise_App_Config()
	if err != nil {
		log.Fatal("error with initialising_App_Config")
	}

	err = server.Create_and_Start_Server(app_ptr)
	if err != nil {
		log.Fatal("error with starting the server")
	}

	//just wait till the program is ended.
	//todo make this cleaner
	if app_ptr.InProduction {
		for {
		} //endless loop
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("End the program by pressing enter")
		_, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
	}

}
