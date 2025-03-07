package main

import (
	"log"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/Lukasloetscher/Go_my_website/pkg/server"
)

func main() {
	var app_ptr *appconfig.AppConfig
	app_ptr, err := appconfig.Initialize_App_Config()
	if err != nil {
		log.Fatal("error with initializing_App_Config")
	}

	err = server.Create_and_Start_Server(app_ptr)
	if err != nil {
		log.Fatal("error with starting the server")
	}

	//just wait till the program is ended.
	//todo make this cleaner
	for {
	} //endless loop
}
