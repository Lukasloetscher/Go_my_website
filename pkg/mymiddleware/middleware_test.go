package mymiddleware

import (
	"testing"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

// TestInclude_Middleware()
// Tests if this code runs without an error
// Does currently not test content.
func TestInclude_Middleware(t *testing.T) {
	t.Run("Include_Middleware", func(t *testing.T) {
		app_ptr, _ := appconfig.Initialise_App_Config()
		mux := chi.NewMux()
		err := Include_Middleware(app_ptr, mux)
		if err != nil {
			t.Error()
		}
	})
}
