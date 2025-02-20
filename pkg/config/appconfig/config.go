package appconfig

// Holds the application config
type AppConfig struct {
	Port         string        //includes the : already!
	InProduction bool          //if this is the productive or dev envirement
	Gen_Pages    Generic_Pages //The data needed to add generic pages.
}

type Generic_Pages struct {
	Webpage_Location string //The location, where the generic files are.
	Layout_Location  string // The position of the Folder whcih contains the Layouts. This needs to be a glob (i.e. .../*.layout.tmpl)
	Add_Navbar       bool   //If the generic pages should include the base navbar (navbar.layout.tmpl)
}
