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
	fmt.Println(dob)
	fmt.Println(dob.DOB)
}
