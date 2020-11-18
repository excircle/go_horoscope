package main

import (
	"fmt"
	"io"
)

func clear() {
	fmt.Print("\033[H\033[2J")
}

//Terminal warning for user
func warnUser(o io.Writer, warning string, blank int) {
	warn := "ERROR: Input is not valid.\n\nYou entered:"

	if blank == 0 {
		clear()
		o.Write([]byte(fmt.Sprintf("%s%s", warn, warning)))
	} else {
		clear()
		o.Write([]byte(fmt.Sprintf("\nERROR: Input is not valid.\n%s\n", warning)))
	}
}
