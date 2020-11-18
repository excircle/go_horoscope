package main

import "fmt"

func clear() {
	fmt.Print("\033[H\033[2J")
}

//Terminal warning for user
func warnUser(warning string, blank int) {
	warn := "ERROR: Input is not valid.\n\nYou entered:"

	if blank == 0 {
		clear()
		fmt.Printf("%s\n\n%s", warn, warning)
	} else {
		clear()
		fmt.Printf("ERROR: Input is not valid.\n%s\n", warning)
	}
}
