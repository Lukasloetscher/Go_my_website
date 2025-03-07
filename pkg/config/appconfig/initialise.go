package appconfig

import (
	"log"
	"os"

	getPort "github.com/Lukasloetscher/Go_my_website/util/get_port"
)

// Initialize_App_Config() sets up the configuration, for the initial version.
// This returns a pointer to the newly defined Appconfig struct.
func Initialize_App_Config() (app_ptr *AppConfig, err error) {
	//This is currently hardcoded, but should be later loaded from a file or through other means.
	var app_config AppConfig
	app_config.InProduction = initialize_In_Production()
	app_config.SecureWebpage = app_config.InProduction //for simplicity. i want to load this from an env file later anyway...
	app_ptr.SecureWebpageCerts.Location_certFile = "todo"
	app_ptr.SecureWebpageCerts.Location_keyFile = "todo"
	app_config.Port = initialize_Port_Number()
	app_config.Channel_Server_Restart = make(chan error)
	app_config.RestartServerWhenShutdown = true
	var Gen_Pages Generic_Pages
	Gen_Pages.Webpage_Location = "webpages"
	Gen_Pages.Layout_Location = "webpages/0_layouts/*.layout.tmpl"
	Gen_Pages.Add_Navbar = true
	app_config.Gen_Pages = Gen_Pages
	return &app_config, nil
}

func initialize_Port_Number() string {
	return ":" + getPort.Get_Port_From_Env()
}

const isProd_env = "PROD"

// initialize_In_Production reads if it is in production from env. if not found returns false
// TODO refractor and combine with Port!
func initialize_In_Production() (result bool) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println("Recovered from error")
			log.Println(r)
			result = false
		}
	}()

	isProd, ok := os.LookupEnv(isProd_env)
	if ok {
		return isTrue(isProd)
	} else {
		return false
	}
}

// isTrue converts string to bool
// todo this should not be here -> refractor
func isTrue(cond string) bool {
	if cond == "true" {
		return true
	} else if cond == "false" {
		return false
	} else {
		panic("incorrect input received isTrue input was neither true nor false.")
	}
}
