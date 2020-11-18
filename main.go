package main

import (
	"fmt"
	"os"
)

var (
	dob DOB
)

func main() {
	//Wipe screen
	clear()
	for dob.Validated != true {
		dob.gatherDOB()
		dob.validateInput(os.Stdout)
	}

	answer := checkSign(dob)

	fmt.Printf("%s\n", answer[dob.IntMonth].Longname)
}
