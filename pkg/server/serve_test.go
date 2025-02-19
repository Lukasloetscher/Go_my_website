package server

import (
	"testing"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

// TestCreate_and_Start_Server() Tests the function Create and Start Server
// Currently only checks if the this goes through without an error.
func TestCreate_and_Start_Server(t *testing.T) {
	t.Run("Create_and_Start_Server", func(t *testing.T) {
		app_ptr, _ := appconfig.Initialise_App_Config()
		err := Create_and_Start_Server(app_ptr)
		if err != nil {
			t.Error()
		}
	})
}

// TestCreate_Server() Tests the function Create Server
// Currently only uses an empty mux
func TestCreate_Server(t *testing.T) {
	t.Run("Create_and_Start_Server", func(t *testing.T) {
		mux := chi.NewRouter()
		app_ptr, _ := appconfig.Initialise_App_Config()
		srv := Create_Server(app_ptr, mux)
		if srv == nil {
			t.Error()
		}
	})
}
