package handlers

import (
	"testing"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

// Testdd_Manual_Handlers()
// Tests if this code runs without an error
// Does currently not test content.
func TestAdd_Manual_Handlers(t *testing.T) {
	t.Run("Include_Middleware", func(t *testing.T) {
		app_ptr, _ := appconfig.Initialise_App_Config()
		mux := chi.NewMux()
		err := Add_Manual_Handlers(app_ptr, mux)
		if err != nil {
			t.Error()
		}
	})
}
