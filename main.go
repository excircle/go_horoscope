package main

import "fmt"

var (
	dob DOB
)

//redirect branch
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
