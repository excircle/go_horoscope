package main

import (
	"os"
	"testing"
)

func TestInvalidDOBInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"12"},
		{"12/14"},
	}
	for _, test := range tests {
		//Create and assign test input
		obj := DOB{}
		obj.DOB = test.input

		//redirect terminal user warning so it doesn't appear in testing output
		rescueStdout := os.Stdout
		_, w, _ := os.Pipe()
		os.Stdout = w

		//Test if test input incorrect sets obj.Validated == true
		//Should be false if working correctly
		obj.validateDOBInput()

		//Set stdout back
		w.Close()
		os.Stdout = rescueStdout

		//Bad inputs should not have set obj.Validated == true
		//FAIL if true
		if obj.Validated == true {
			t.Errorf("validateDOBInput() = have: %s & 'true', want: false ", test.input)
		}
	}
}
