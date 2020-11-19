package main

import (
	"fmt"
	"os"
)

var (
	dob DOB
	z   zodiak
)

func main() {
	//Wipe screen
	clear()
	for dob.Validated != true {
		dob.gatherDOB()
		dob.validateInput(os.Stdout)
	}

	z = buildZodiak()
	answer := checkSign(dob)

	fmt.Printf("Your sign is: %s\n", answer)
}
