package main

import "fmt"

var (
	dob DOB
)

func main() {
	//Wipe screen
	Clear()
	dob.GatherDOB()
	fmt.Println(dob.DOB)

}
