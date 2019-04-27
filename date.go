package mantis

import (
	"regexp"
	"time"
)

// Date struct
type Date struct {
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Hour       string `json:"hour"`
	Minute     string `json:"minute"`
	Second     string `json:"second"`
	Nanosecond string `json:"nanosecond"`
}

// Takes a date string and returns a Date struct
func StringToDate(date string) Date {
	var dateReturn Date

	if date == "" {
		return Date{}
	}

	pattern := `(\d\d\d\d)-(\d\d)-(\d\d)\s(\d\d)\:(\d\d):(\d\d)`
	if matched, err := regexp.MatchString(pattern, date); err == nil && matched == true {
		// We have a 0000-00-00 00:00:00 format date
		dateReturn = Date{
			Year:       string([]rune(date)[0:4]),
			Month:      string([]rune(date)[5:7]),
			Day:        string([]rune(date)[8:10]),
			Hour:       string([]rune(date)[11:13]),
			Minute:     string([]rune(date)[14:16]),
			Second:     string([]rune(date)[17:19]),
			Nanosecond: string([]rune(date)[20:23]),
		}
	} else {
		t, er := time.Parse("2006-01-02T15:04:05.000Z", date)
		HandleError(er)
		dateReturn = Date{
			Year:       string(t.Year()),
			Month:      string(t.Month()),
			Day:        string(t.Day()),
			Hour:       string(t.Hour()),
			Minute:     string(t.Minute()),
			Second:     string(t.Second()),
			Nanosecond: string(t.Nanosecond()),
		}
	}

	//fmt.Println(dateReturn)
	return dateReturn
}

// Takes a Date struct and returns a string
func (d *Date) DateToString() string {
	if d.Year == "" {
		return ""
	}

	dateString := d.Year + "-" + d.Month + "-" + d.Day + " " + d.Hour + ":" + d.Hour + ":" + d.Minute

	return dateString
}
