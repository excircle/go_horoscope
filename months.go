package main

type Month struct {
	Longname  string
	Shortname string
}

var monthArray = map[int]Month{
	1:  {Longname: "Janurary", Shortname: "Jan"},
	2:  {Longname: "February", Shortname: "Feb"},
	3:  {Longname: "March", Shortname: "Mar"},
	4:  {Longname: "April", Shortname: "Apr"},
	5:  {Longname: "May", Shortname: "May"},
	6:  {Longname: "June", Shortname: "Jun"},
	7:  {Longname: "July", Shortname: "Jul"},
	8:  {Longname: "August", Shortname: "Aug"},
	9:  {Longname: "September", Shortname: "Sep"},
	10: {Longname: "October", Shortname: "Oct"},
	11: {Longname: "November", Shortname: "Nov"},
	12: {Longname: "December", Shortname: "Dec"}}
