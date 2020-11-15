package main

import (
	"os"
	"testing"
)

//Ensure code catches bad input & doesn't set obj.Validate true
func TestInvalidInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		//too short
		{"12"},
		{"12/14"},
		//bad month
		{"13/14/1989"},
		//bad day
		{"12/32/1989"},
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
		obj.validateInput()

		//Set stdout back
		w.Close()
		os.Stdout = rescueStdout

		//Bad inputs should not have set obj.Validated == true
		//FAIL if true
		if obj.Validated == true {
			t.Errorf("validateInput() = have: %s & '%v', want: false ", test.input, obj.Validated)
		}
	}
}

//Test handling short input
func TestShortInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		//Single digit month
		{"6/14/89"},
		//Single digit day
		{"06/6/89"},
		//Year is 2 digits instead of 4
		{"12/14/89"},
	}
	for _, test := range tests {
		//Create and assign test input
		obj := DOB{}
		obj.DOB = test.input

		//redirect terminal user warning so it doesn't appear in testing output
		rescueStdout := os.Stdout
		_, w, _ := os.Pipe()
		os.Stdout = w

		//Test if test input incorrectly leaves obj.Validated == false
		//Should be true if working correctly
		obj.validateInput()

		//Set stdout back
		w.Close()
		os.Stdout = rescueStdout

		//Bad inputs should not have set obj.Validated == true
		//FAIL if true
		if obj.Validated == false {
			t.Errorf("validateInput() = have: %s & '%v', want: true ", test.input, obj.Validated)
			t.Errorf("obj.DOB: %s", obj.DOB)
		}
	}
}

func TestValidInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"12/14/1989"},
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
		obj.validateInput()

		//Set stdout back
		w.Close()
		os.Stdout = rescueStdout

		//Bad inputs should not have set obj.Validated == true
		//FAIL if true
		if obj.Validated != true {
			t.Errorf("validateInput() = have: %s & '%v', want: true ", test.input, obj.Validated)
			t.Errorf("%v", obj)
		}
	}
}
