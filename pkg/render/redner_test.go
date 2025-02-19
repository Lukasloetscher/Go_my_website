package render

import (
	"fmt"
	"testing"

	my_testing "github.com/Lukasloetscher/Go_my_website/cmd"
)

// TestCheck_If_Layout_Exists()
// Tests if this code runs without an error
// TODO look up if there is a way, to mok a filesystem, for testing.
func TestCheck_If_Layout_Exists(t *testing.T) {
	//test data
	var tests = []struct {
		Layout_Path string
		ok          bool
		panic       bool
	}{
		{"", false, false},
		{"ghksg", false, true}, //incorrect string
		{"fagf", false, true},
		{"webpages/layouts/", true, false},
	}

	//test functionality
	for _, current_test := range tests {
		testname := "Check_If_Layout_Exists tested with Layout_Path: " + current_test.Layout_Path
		t.Run(testname, func(t *testing.T) {

			defer my_testing.Check_For_Panic(t, current_test.panic)

			if current_test.ok != Check_If_Layout_Exists(current_test.Layout_Path) {
				fmt.Print("Incorrect Result")
				t.Error()
			}

		})
	}
}

// TestLoad_Template_From_File()
// Tests if this code runs without an error
// TODO look up if there is a way, to mok a filesystem, for testing.
// also currecntly this code only tests if the code runs
func TestLoad_Template_From_File(t *testing.T) {
	t.Run("TestLoad_Template_From_File", func(t *testing.T) {
		defer my_testing.Check_For_Panic(t, false)

		var paths Render_Filepaths
		paths.Filepath_Html = ""
		paths.Filepath_Layout = ""
		ts := Load_Template_From_File(paths)
		if ts == nil {
			t.Error()
		}
	})
}
