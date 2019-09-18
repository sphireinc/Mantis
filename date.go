package mantis

import (
	"fmt"
	"time"
)

// Date struct
type Date struct {
	Year       int          `json:"year"`
	Month      time.Month   `json:"month"`
	Day        int          `json:"day"`
	Hour       int          `json:"hour"`
	Minute     int          `json:"minute"`
	Second     int          `json:"second"`
	Nanosecond int          `json:"nanosecond"`
	Unix       int64        `json:"unix"`
	WeekDay    time.Weekday `json:"week_day"`
	YearDay    int          `json:"year_day"`
}

// CurrentTime Returns a the current date
func CurrentTime() Date {
	current := time.Now()

	return Date{
		Year:       current.Year(),
		Month:      current.Month(),
		Day:        current.Day(),
		Hour:       current.Hour(),
		Nanosecond: current.Nanosecond(),
		Second:     current.Second(),
		Minute:     current.Minute(),
		Unix:       current.Unix(),
		YearDay:    current.YearDay(),
		WeekDay:    current.Weekday(),
	}
}

// Takes a date string YYYY-MM-DD HH:MM:SS and returns a Date struct
func StringToDate(date string) Date {
	if date == "" {
		return Date{}
	}

	t, err := time.Parse("2006-01-02T15:04:05.000Z", date)
	HandleError("Error in StringToDate time.Parse", err)

	return Date{
		Year:   t.Year(),
		Month:  t.Month(),
		Day:    t.Day(),
		Hour:   t.Hour(),
		Minute: t.Minute(),
		Second: t.Second(),
	}
}

// DateToString Takes a Date struct and returns a string in format YYYY-MM-DD HH:II:SS
func (d *Date) DateToString() string {
	return fmt.Sprintf("%s-%s-%s %s:%s:%s", itos(d.Year), itos(int(d.Month)), itos(d.Day), itos(d.Hour), itos(d.Minute), itos(d.Second))
}

// itos Converts an int to a string, prepends zero if less than 10
func itos(intVal int) string {
	if intVal == 0 {
		return "0"
	}
	intValStr := string(intVal)
	if intVal < 10 {
		return "0" + intValStr
	}
	return intValStr
}
