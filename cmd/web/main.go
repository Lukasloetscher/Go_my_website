package main

import (
	"log"
	"os"

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

	//todo this should be refactored into better read able code.
	var restart_count int
	restart_count = 0
	for {
		select {
		case Server_did_shut_down := <-app_ptr.Channel_Server_Restart:
			log.Println("Received Signal that Server shut down.")
			log.Println(Server_did_shut_down)
			if app_ptr.RestartServerWhenShutdown {
				restart_count += 1
				log.Println("restart_count = ", restart_count)
				if restart_count > 5 && app_ptr.SecureWebpage {
					log.Println("switching from https to http due to too many failed starts.")
					app_ptr.SecureWebpage = false
				}
				if restart_count > 10 { //we should reset this variable after some time...
					log.Println("Exiting program after too many restarts attempts.")
					os.Exit(-1)
				}

				log.Println("restarting server")
				err = server.Create_and_Start_Server(app_ptr)
				if err != nil {
					log.Fatal("error with starting the server")
				}
			} else {
				log.Println("Terminating program")
				os.Exit(-1)
			}
		}
	}

}
