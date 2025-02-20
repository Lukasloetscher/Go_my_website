package appconfig

import getport "github.com/Lukasloetscher/Go_my_website/util/get_port"

//Initialise_App_Config() sets up the configuration, for the initial version.
//This returns a pointer to the newly defined Appconfig struct.
func Initialise_App_Config() (app_ptr *AppConfig, err error) {
	//This is currently hardcoded, but should be later loaded from a file or through other means.
	var app_config AppConfig
	app_config.InProduction = false
	app_config.Port = initialse_Port_Number()
	var Gen_Pages Generic_Pages
	Gen_Pages.Webpage_Location = "webpages"
	Gen_Pages.Layout_Location = "webpages/0_layouts/*.layout.tmpl"
	Gen_Pages.Add_Navbar = true
	app_config.Gen_Pages = Gen_Pages
	return &app_config, nil
}

func initialse_Port_Number() string {
	return ":" + getport.Get_Port_From_Env()
}
