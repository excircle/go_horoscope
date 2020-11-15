package main

import (
	"fmt"
	"strings"
)

type DOB struct {
	DOB       string
	StrDOB    []string
	IntDOB    []int
	StrMonth  string
	StrDay    string
	StrYear   string
	IntMonth  int
	IntDay    int
	IntYear   int
	Validated bool
}

func (d *DOB) gatherDOB() {
	fmt.Println("Please enter your birthday (mm/dd/yyyy)")
	fmt.Scanln(&dob.DOB)
}

//func (d *DOB) create

func (d *DOB) validateDOBInput() {
	d.StrDOB = strings.Split(d.DOB, "/")

	if len(d.StrDOB) != 3 {
		WarnUser(d.DOB+"\n\n", 0)
		return
	}
	d.Validated = true
}
