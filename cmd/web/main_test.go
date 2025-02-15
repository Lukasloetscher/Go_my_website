package main

import (
	"testing"
)

func TestCheck_If_Valid_Port(t *testing.T) {
	var tests = []struct {
		port   string
		result bool
	}{
		{"8080", true},
		{"1215a125", false},
		{"fagf", false},
		{"1234", true},
	}
	for _, current_test := range tests {
		testname := "Port: " + current_test.port
		t.Run(testname, func(t *testing.T) {
			if current_test.result != Check_If_Valid_Port(current_test.port) {
				t.Error()
			}
		})
	}

}
