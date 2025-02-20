package htmlerrors

import (
	"log"
	"net/http"
)

// Catch_Panic_And_Report_Internal_Error
// Catsches errors and writes a error message to the frontend
func Catch_Panic_And_Report_Internal_Error(w http.ResponseWriter) {
	r := recover()
	if r != nil {
		log.Println("Recovcered from Panic")
		log.Println(r)
		w.Write([]byte("405 Internal error"))
	}
}
