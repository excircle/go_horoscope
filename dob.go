package main

import "fmt"

type DOB struct {
	DOB      string
	StrDOB   []string
	IntDOB   []int
	StrMonth string
	StrDay   string
	StrYear  string
	IntMonth int
	IntDay   int
	IntYear  int
}

func (d *DOB) SetDOB(s string) {
	d.DOB = s
}

func (d *DOB) SayDOB() string {
	return d.DOB
}

func (d *DOB) GatherDOB() {
	fmt.Println("Please enter your birthday (mm/dd/yyyy)")
	fmt.Scanln(&dob.DOB)
}

func (d *DOB) validateDOBInput() {

}
