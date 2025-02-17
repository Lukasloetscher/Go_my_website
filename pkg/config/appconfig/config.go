package appconfig

// Holds the application config
type AppConfig struct {
	Port         string //includes the : already!
	InProduction bool   //if this is the productive or dev envirement
}
