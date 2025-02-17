package appconfig

import "testing"

//TestInitialise_App_Config()
//Tests if a valid pointer is created.
//Does currently not test content.
func TestInitialise_App_Config(t *testing.T) {
	t.Run("Initialise App_Config", func(t *testing.T) {
		_, err := Initialise_App_Config()
		if err != nil {
			t.Error()
		}
	})
}

//Testinitialse_Port_Number()
//Makes sure the string is :8080
func Testinitialse_Port_Number(t *testing.T) {
	t.Run("Initialise App_Config", func(t *testing.T) {
		Port := initialse_Port_Number()
		if Port != ":8080" {
			t.Error()
		}
	})
}
