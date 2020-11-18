package main

import (
	"fmt"
	"log"
	"strconv"
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

func (d *DOB) createData() {
	//Create individual string data
	d.StrMonth, d.StrDay, d.StrYear = d.StrDOB[0], d.StrDOB[1], d.StrDOB[2]

	//Create int data array
	for k := range d.StrDOB {
		newint, err := strconv.Atoi(d.StrDOB[k])
		if err != nil {
			log.Println(err)
		}
		d.IntDOB = append(d.IntDOB, newint)
	}

	//Create individual int data
	d.IntMonth, d.IntDay, d.IntYear = d.IntDOB[0], d.IntDOB[1], d.IntDOB[2]
}

func (d *DOB) checkData() bool {
	//Check if month is within Jan-Dec/1-12 range
	if d.IntMonth < 1 || d.IntMonth > 12 {
		warnUser("Wrong month -> "+d.DOB+"\n\n", 0)
		return false
	}
	//Check if day within 1-31 range
	if d.IntDay < 1 || d.IntDay > 31 {
		warnUser("Wrong day -> "+d.DOB+"\n\n", 0)
		return false
	}

	//Fix Month if not in 'mm' format
	if len(d.StrMonth) < 2 {
		if d.IntMonth >= 1 && d.IntMonth <= 9 {
			d.StrMonth = "0" + d.StrMonth
			d.DOB = fmt.Sprintf("%s/%s/%s", d.StrMonth, d.StrDay, d.StrYear)
		}
	}

	//Fix Day if not in 'dd' format
	if len(d.StrDay) < 2 {
		if d.IntDay >= 1 && d.IntDay <= 9 {
			d.StrDay = "0" + d.StrDay
			d.DOB = fmt.Sprintf("%s/%s/%s", d.StrMonth, d.StrDay, d.StrYear)
			d.createData()
		}
	}

	//Check if 4-digit year supplied
	if len(d.StrYear) < 4 {
		//if not 4-digit, check if it's between 1-99 and modify for usage (Example: '76' modifies into '1976')
		if d.IntYear >= 1 && d.IntYear <= 99 {
			d.IntYear += 1900
			d.DOB = fmt.Sprintf("%s/%s/%s", d.StrMonth, d.StrDay, strconv.Itoa(d.IntYear))
			d.StrDOB = strings.Split(d.DOB, "/")
			d.createData()
			return true
		} else {
			return false
		}
	}
	return true
}
func (d *DOB) validateInput() {
	d.StrDOB = strings.Split(d.DOB, "/")

	//Reject invalidate date input
	if len(d.StrDOB) != 3 {
		warnUser("\nIncorrect format. Use slashes!\n", 1)
		return
	}

	//Create day, month, and year as individual Int & Str fields
	d.createData()
	if d.checkData() == false {
		return
	}

	d.Validated = true
}
