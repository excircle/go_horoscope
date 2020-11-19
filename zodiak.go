package main

import "time"

type sign struct {
	Name  string
	Start time.Time
	End   time.Time
}

type month struct {
	Longname  string
	Shortname string
	Signs     []sign
}

type zodiak struct {
	Zodiak map[int]month
}

func buildZodiak() zodiak {
	var x zodiak = zodiak{
		map[int]month{
			//month_id|Longname|Shortname|[]sign|Name|time.Time|time.Time
			1:  {"Janurary", "Jan", []sign{{"Capricorn", getDate(dob.IntYear, 1, 1), getDate(dob.IntYear, 1, 20)}, {"Aquarius", getDate(dob.IntYear, 1, 21), getDate(dob.IntYear, 2, 1)}}},
			2:  {"February", "Feb", []sign{{"Aquarius", getDate(dob.IntYear, 2, 1), getDate(dob.IntYear, 2, 18)}, {"Pisces", getDate(dob.IntYear, 2, 19), getDate(dob.IntYear, 3, 1)}}},
			3:  {"March", "Mar", []sign{{"Pisces", getDate(dob.IntYear, 2, 19), getDate(dob.IntYear, 3, 1)}, {"Aries", getDate(dob.IntYear, 3, 21), getDate(dob.IntYear, 4, 1)}}},
			4:  {"April", "Apr", []sign{{"Aries", getDate(dob.IntYear, 4, 1), getDate(dob.IntYear, 4, 20)}, {"Taurus", getDate(dob.IntYear, 4, 21), getDate(dob.IntYear, 5, 1)}}},
			5:  {"May", "May", []sign{{"Taurus", getDate(dob.IntYear, 5, 1), getDate(dob.IntYear, 5, 21)}, {"Gemini", getDate(dob.IntYear, 5, 22), getDate(dob.IntYear, 6, 1)}}},
			6:  {"June", "Jun", []sign{{"Gemini", getDate(dob.IntYear, 6, 1), getDate(dob.IntYear, 6, 21)}, {"Cancer", getDate(dob.IntYear, 6, 22), getDate(dob.IntYear, 7, 1)}}},
			7:  {"July", "Jul", []sign{{"Cancer", getDate(dob.IntYear, 7, 1), getDate(dob.IntYear, 7, 22)}, {"Leo", getDate(dob.IntYear, 7, 23), getDate(dob.IntYear, 8, 1)}}},
			8:  {"August", "Aug", []sign{{"Leo", getDate(dob.IntYear, 8, 1), getDate(dob.IntYear, 8, 23)}, {"Virgo", getDate(dob.IntYear, 8, 24), getDate(dob.IntYear, 9, 1)}}},
			9:  {"September", "Sep", []sign{{"Virgo", getDate(dob.IntYear, 9, 1), getDate(dob.IntYear, 9, 23)}, {"Libra", getDate(dob.IntYear, 9, 24), getDate(dob.IntYear, 10, 1)}}},
			10: {"October", "Oct", []sign{{"Libra", getDate(dob.IntYear, 10, 1), getDate(dob.IntYear, 10, 23)}, {"Scorpio", getDate(dob.IntYear, 10, 24), getDate(dob.IntYear, 11, 1)}}},
			11: {"November", "Nov", []sign{{"Scorpio", getDate(dob.IntYear, 11, 1), getDate(dob.IntYear, 11, 22)}, {"Sagittarius", getDate(dob.IntYear, 11, 23), getDate(dob.IntYear, 12, 1)}}},
			12: {"December", "Dec", []sign{{"Sagittarius", getDate(dob.IntYear, 12, 1), getDate(dob.IntYear, 12, 21)}, {"Capricorn", getDate(dob.IntYear, 12, 22), getDate(dob.IntYear, 12, 31)}}}}}

	return x
}
