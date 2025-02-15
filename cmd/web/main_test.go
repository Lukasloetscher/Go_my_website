package main

import (
	"fmt"
	"testing"
)

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
			if current_test.result != Check_If_Valid_Port(current_test.port) {
				t.Error()
			}
		})
	}

}
func TestGet_Port_From_Env(t *testing.T) {
	//test data
	var tests = []struct {
		port string
		ok   bool
	}{
		{"8080", true},
		{"1215a125", false},
		{"fagf", false},
		{"1234", true},
		{"", true},
	}

	//test functionality
	for _, current_test := range tests {
		testname := "Get_Port_From_Env tested with Port: " + current_test.port
		t.Run(testname, func(t *testing.T) {
			if current_test.port != "" { //We also want to test the case, where no ev, was defined.
				t.Setenv("PORT", current_test.port)
			}
			//test for cases here the code panics
			defer func() {
				r := recover()
				did_panic := r != nil
				if did_panic && current_test.ok {
					fmt.Print("did incorrectly panic")
					t.Error()
				}
				if !did_panic && !current_test.ok {
					fmt.Print("did not panic but should have")
					t.Error()
				}
			}()

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
