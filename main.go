package main

import "fmt"

var (
	dob DOB
)

func main() {
	//Wipe screen
	Clear()
	for dob.Validated != true {
		dob.gatherDOB()
		dob.validateInput()
	}
	fmt.Println(dob)
	fmt.Println(dob.DOB)
}
