package main

func checkSign(d DOB) string {
	date := getDate(d.IntYear, d.IntMonth, d.IntDay)
	for _, v := range z.Zodiak[d.IntMonth].Signs {
		a := date.After(v.Start)
		b := date.Before(v.End)
		if a == true && b == true {
			return v.Name
		}
	}
	return "ERROR: Sign not found!"
}
