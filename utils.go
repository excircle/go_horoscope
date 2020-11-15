package main

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J")
}

//Terminal warning for user
func WarnUser(warning string, blank int) {
	warn := "ERROR: Input is not valid.\n\nYou entered:"

	if blank == 0 {
		Clear()
		fmt.Printf("%s\n\n%s", warn, warning)
	} else {
		Clear()
		fmt.Printf("ERROR: Input is not valid.\n%s\n", warning)
	}
}
