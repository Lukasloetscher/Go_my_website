package generichandlers

import (
	"testing"

	"github.com/Lukasloetscher/Go_my_website/pkg/config/appconfig"
	"github.com/go-chi/chi"
)

// TestAdd_Generic_Handlers()
// Tests if this code runs without an error
// Does currently not test content.
func TestAdd_Generic_Handlers(t *testing.T) {
	t.Run("Include_Middleware", func(t *testing.T) {
		app_ptr, _ := appconfig.Initialize_App_Config()
		mux := chi.NewMux()
		err := Add_Generic_Handlers(app_ptr, mux)
		if err != nil {
			t.Error()
		}
	})
}
