package appconfig

import "testing"

//Test_Initialize_App_Config()
//Tests if a valid pointer is created.
//Does currently not test content.
func Test_Initialize_App_Config(t *testing.T) {
	t.Run("Initialize App_Config", func(t *testing.T) {
		_, err := Initialize_App_Config()
		if err != nil {
			t.Error()
		}
	})
}

//Test_initialize_Port_Number()
//Makes sure the string is :8080
func Test_initialize_Port_Number(t *testing.T) {
	t.Run("initialize_Port_Number", func(t *testing.T) {
		Port := initialize_Port_Number()
		if Port != ":8080" {
			t.Error()
		}
	})
}
