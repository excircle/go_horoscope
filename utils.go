package main

import (
	"fmt"
	"io"
	"time"
)

//Clear terminal screen
func clear() {
	fmt.Print("\033[H\033[2J")
}

//Get time.Time object from 3 int argument
func getDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
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
