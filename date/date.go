package date

import (
	"fmt"
	mantisError "github.com/sphireinc/mantis/error"
	"strconv"
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

// CurrentTime is a Date factory
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

// StringToDate takes a date string YYYY-MM-DD HH:MM:SS and returns a Date struct
func StringToDate(date string) Date {
	if date == "" {
		return Date{}
	}

	current, err := time.Parse("2006-01-02 15:04:05", date)
	mantisError.HandleError("Error in StringToDate time.Parse", err)

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

// DateToString takes a Date struct and returns a string in format YYYY-MM-DD HH:II:SS
func (d *Date) DateToString() string {
	return fmt.Sprintf("%s-%s-%s %s:%s:%s", itos(d.Year), itos(int(d.Month)), itos(d.Day), itos(d.Hour), itos(d.Minute), itos(d.Second))
}

// itos converts an int to a string, prepends zero if less than 10
func itos(intVal int) string {
	if intVal == 0 {
		return "00"
	}
	intValStr := strconv.Itoa(intVal)
	if intVal < 10 {
		return "0" + intValStr
	}
	return intValStr
}
