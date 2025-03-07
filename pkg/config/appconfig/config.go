package appconfig

// Holds the application config
type AppConfig struct {
	Port                      string               //includes the : already!
	InProduction              bool                 //if this is the productive or dev environment
	SecureWebpage             bool                 //false = http true = https
	SecureWebpageCerts        Certificate_location //file location of the certs used for https
	RestartServerWhenShutdown bool                 //false = no restarts, true restarts.
	Gen_Pages                 Generic_Pages        //The data needed to add generic pages.
	Channel_Server_Restart    chan error           //used to send signals when the server shuts down.
}

type Generic_Pages struct {
	Webpage_Location string //The location, where the generic files are.
	Layout_Location  string // The position of the Folder which  contains the Layouts. This needs to be a glob (i.e. .../*.layout.tmpl)
	Add_Navbar       bool   //If the generic pages should include the base navbar (navbar.layout.tmpl)
}

type Certificate_location struct {
	Location_certFile string
	Location_keyFile  string
}
