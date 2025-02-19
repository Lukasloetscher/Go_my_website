package getport

import (
	"fmt"
	"testing"

	my_testing "github.com/Lukasloetscher/Go_my_website/cmd"
)

// TestCheck_If_Valid_Port()
// Test to check if the port validation works as intended
func TestCheck_If_Valid_Port(t *testing.T) {
	//test data
	var tests = []struct {
		port   string
		result bool
	}{
		{"8080", true},
		{"1215a125", false},
		{"fagf", false},
		{"1234", true},
		{"", false},
	}

	//test functionality
	for _, current_test := range tests {
		testname := "Check_If_Valid_Port tested with Port: " + current_test.port
		t.Run(testname, func(t *testing.T) {
			if current_test.result != check_If_Valid_Port(current_test.port) {
				t.Error()
			}
		})
	}

}

// TestGet_Port_From_Env()
// Test to check if the port is correctly read from the envirement.
func TestGet_Port_From_Env(t *testing.T) {
	//test data
	var tests = []struct {
		port  string
		ok    bool
		panic bool
	}{
		{"8080", true, false},
		{"1215a125", false, true},
		{"fagf", false, true},
		{"1234", true, false},
		{"", true, false},
	}

	//test functionality
	for _, current_test := range tests {
		testname := "Get_Port_From_Env tested with Port: " + current_test.port
		t.Run(testname, func(t *testing.T) {
			if current_test.port != "" { //We also want to test the case, where no env, was defined.
				t.Setenv("PORT", current_test.port)
			}
			//test for cases here the code panics
			defer my_testing.Check_For_Panic(t, current_test.panic)

			port_result := Get_Port_From_Env()
			if port_result != current_test.port && !current_test.ok { //special ase for when no env variable is set
				if !(port_result == "8080" && current_test.port == "") {
					fmt.Print("returned incorrect port number")
					t.Error()
				}
			}

		})
	}
}
