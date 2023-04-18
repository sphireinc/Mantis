package date

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (d *Date) String() string {
	marshaledStruct, err := json.Marshal(d)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
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
func StringToDate(date string) (Date, error) {
	if date == "" {
		return Date{}, errors.New("no date string given")
	}

	current, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return Date{}, err
	}

	dateObj := Date{
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
	return dateObj, nil
}

// DateToString takes a Date struct and returns a string in format YYYY-MM-DD HH:II:SS
//func (d *Date) DateToString() string {
//	return fmt.Sprintf("%s-%s-%s %s:%s:%s", itos(d.Year), itos(int(d.Month)), itos(d.Day), itos(d.Hour), itos(d.Minute), itos(d.Second))
//}

// DateToString converts a Date object to a string representation in a given format
func (d *Date) DateToString(format string) (string, error) {
	t := time.Date(d.Year, d.Month, d.Day, d.Hour, d.Minute, d.Second, d.Nanosecond, time.UTC)
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return t.Format(format), nil
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

// NewDate creates a new Date object from a given year, month, day, hour, minute, second, and nanosecond
func NewDate(year, month, day, hour, minute, second, nanosecond int) (*Date, error) {
	t := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.UTC)
	d := &Date{
		Year:       t.Year(),
		Month:      t.Month(),
		Day:        t.Day(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
		Nanosecond: t.Nanosecond(),
		Unix:       t.Unix(),
		WeekDay:    t.Weekday(),
		YearDay:    t.YearDay(),
	}

	if d.Year != year || int(d.Month) != month || d.Day != day ||
		d.Hour != hour || d.Minute != minute || d.Second != second || d.Nanosecond != nanosecond {
		return nil, fmt.Errorf("invalid date and time")
	}

	return d, nil
}

// ParseDate parses a string representation of a date in a given format and returns a Date object
func ParseDate(dateStr, format string) (*Date, error) {
	t, err := time.Parse(format, dateStr)
	if err != nil {
		return nil, err
	}
	d := &Date{
		Year:       t.Year(),
		Month:      t.Month(),
		Day:        t.Day(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
		Nanosecond: t.Nanosecond(),
		Unix:       t.Unix(),
		WeekDay:    t.Weekday(),
		YearDay:    t.YearDay(),
	}
	return d, nil
}

// GetWeekDayName returns the name of the week day (e.g. "Monday", "Tuesday") for a given Date object
func GetWeekDayName(date *Date) string {
	return date.WeekDay.String()
}

// AddDays adds a given number of days to a Date object and returns a new Date object
func AddDays(date *Date, days int) *Date {
	t := time.Date(date.Year, date.Month, date.Day, date.Hour, date.Minute, date.Second, date.Nanosecond, time.UTC)
	t = t.AddDate(0, 0, days)
	d := &Date{
		Year:       t.Year(),
		Month:      t.Month(),
		Day:        t.Day(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
		Nanosecond: t.Nanosecond(),
		Unix:       t.Unix(),
		WeekDay:    t.Weekday(),
		YearDay:    t.YearDay(),
	}
	return d
}

// SubtractDates subtracts two Date objects and returns the duration between them in seconds or as a time.Duration object
func SubtractDates(date1, date2 *Date) (int64, time.Duration) {
	t1 := time.Date(date1.Year, date1.Month, date1.Day, date1.Hour, date1.Minute, date1.Second, date1.Nanosecond, time.UTC)
	t2 := time.Date(date2.Year, date2.Month, date2.Day, date2.Hour, date2.Minute, date2.Second, date2.Nanosecond, time.UTC)

	diff := t1.Sub(t2)
	return int64(diff.Seconds()), diff
}

// IsLeapYear checks whether a given year is a leap year or not and returns a boolean value
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns the number of days in a given month and year
func DaysInMonth(year, month int) (int, error) {
	if month < 1 || month > 12 {
		return 0, fmt.Errorf("invalid month")
	}

	if month == 2 {
		if IsLeapYear(year) {
			return 29, nil
		}
		return 28, nil
	}

	return time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Day(), nil
}
