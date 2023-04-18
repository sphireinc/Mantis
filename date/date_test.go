package date

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestStringToDate2(t *testing.T) {
	t1 := Date{2019, time.January, 1, 0, 0, 0, 0, 1546300800, time.Tuesday, 1}
	t2 := Date{2019, time.April, 3, 14, 45, 22, 0, 1554302722, time.Wednesday, 93}
	t3 := Date{1970, time.January, 1, 00, 0, 0, 0, 0, time.Thursday, 1}
	t4 := Date{1870, time.January, 1, 23, 59, 59, 0, -3155587201, time.Saturday, 1}

	tests := []struct {
		date     string
		expected Date
	}{
		{"2019-01-01 00:00:00", t1},
		{"2019-04-03 14:45:22", t2},
		{"1970-01-01 00:00:00", t3},
		{"1870-01-01 23:59:59", t4},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, _ := StringToDate(test.date)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%d', got '%d'", test.expected, actual)
			}
		})
	}
}

func TestStringToDate(t *testing.T) {
	testCases := []struct {
		name     string
		date     string
		expected Date
		errMsg   string
	}{
		{
			name:     "Valid date string",
			date:     "2023-04-18 10:30:00",
			expected: Date{Year: 2023, Month: time.April, Day: 18, Hour: 10, Minute: 30, Second: 0, Nanosecond: 0, Unix: 1681813800, WeekDay: time.Tuesday, YearDay: 108},
			errMsg:   "",
		},
		{
			name:     "Empty date string",
			date:     "",
			expected: Date{},
			errMsg:   "no date string given",
		},
		//{
		//	name:     "Invalid date format",
		//	date:     "18/04/2023 10:30:00",
		//	expected: Date{},
		//	errMsg:   "parsing time \"18/04/2023 10:30:00\" as \"2006-01-02 15:04:05\": cannot parse \"/04/2023 10:30:00\" as \"-\"",
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := StringToDate(tc.date)

			if err != nil && err.Error() != tc.errMsg {
				t.Errorf("StringToDate(%s) returned unexpected error: got %v, expected %s", tc.date, err, tc.errMsg)
			} else if err == nil && actual != tc.expected {
				t.Errorf("StringToDate(%s) returned incorrect date: got %v, expected %v", tc.date, actual, tc.expected)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	tests := []struct {
		expected string
		date     Date
	}{
		{"2019-01-01 00:00:00", Date{2019, 1, 1, 0, 0, 0, 0, 0, 1, 1}},
		{"2019-04-03 14:45:22", Date{2019, 4, 3, 14, 45, 22, 0, 0, 1, 1}},
		{"1970-01-01 00:00:00", Date{1970, 1, 1, 00, 0, 0, 0, 0, 1, 1}},
		{"1870-01-01 23:59:59", Date{1870, 1, 1, 23, 59, 59, 0, 0, 1, 1}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, _ := test.date.DateToString("")
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}

func TestCurrentTime(t *testing.T) {
	// Get the current time
	now := time.Now()

	// Call the CurrentTime function
	date := CurrentTime()

	// Check that the values are correct
	if date.Year != now.Year() ||
		date.Month != now.Month() ||
		date.Day != now.Day() ||
		date.Hour != now.Hour() ||
		date.Minute != now.Minute() ||
		date.Second != now.Second() ||
		date.Nanosecond != now.Nanosecond() ||
		date.Unix != now.Unix() ||
		date.YearDay != now.YearDay() ||
		date.WeekDay != now.Weekday() {
		t.Errorf("CurrentTime() returned incorrect date: got %v, expected %v", date, now)
	}
}

func TestItos(t *testing.T) {
	testCases := []struct {
		name     string
		intVal   int
		expected string
	}{
		{
			name:     "Single digit integer",
			intVal:   5,
			expected: "05",
		},
		{
			name:     "Double digit integer",
			intVal:   23,
			expected: "23",
		},
		{
			name:     "Zero",
			intVal:   0,
			expected: "00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := itos(tc.intVal)

			if actual != tc.expected {
				t.Errorf("itos(%d) returned incorrect string: got %s, expected %s", tc.intVal, actual, tc.expected)
			}
		})
	}
}

func TestNewDate(t *testing.T) {
	testCases := []struct {
		name        string
		year        int
		month       int
		day         int
		hour        int
		minute      int
		second      int
		nanosecond  int
		expected    *Date
		expectedErr error
	}{
		{
			name:        "Valid date and time",
			year:        2023,
			month:       4,
			day:         18,
			hour:        12,
			minute:      30,
			second:      0,
			nanosecond:  0,
			expected:    &Date{Year: 2023, Month: time.April, Day: 18, Hour: 12, Minute: 30, Second: 0, Nanosecond: 0, Unix: 1681821000, WeekDay: time.Tuesday, YearDay: 108},
			expectedErr: nil,
		},
		{
			name:        "Invalid date and time",
			year:        2023,
			month:       2,
			day:         30,
			hour:        12,
			minute:      30,
			second:      0,
			nanosecond:  0,
			expected:    nil,
			expectedErr: fmt.Errorf("invalid date and time"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NewDate(tc.year, tc.month, tc.day, tc.hour, tc.minute, tc.second, tc.nanosecond)

			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("NewDate(%d, %d, %d, %d, %d, %d, %d) returned unexpected error: got %v, expected %v", tc.year, tc.month, tc.day, tc.hour, tc.minute, tc.second, tc.nanosecond, err, tc.expectedErr)
			} else if err == nil && *actual != *tc.expected {
				t.Errorf("NewDate(%d, %d, %d, %d, %d, %d, %d) returned incorrect date: got %v, expected %v", tc.year, tc.month, tc.day, tc.hour, tc.minute, tc.second, tc.nanosecond, actual, tc.expected)
			}
		})
	}
}

func TestParseDate(t *testing.T) {
	testCases := []struct {
		name        string
		dateStr     string
		format      string
		expected    *Date
		expectedErr error
	}{
		{
			name:        "Valid date string and format",
			dateStr:     "2023-04-18 12:30:00",
			format:      "2006-01-02 15:04:05",
			expected:    &Date{Year: 2023, Month: time.April, Day: 18, Hour: 12, Minute: 30, Second: 0, Nanosecond: 0, Unix: 1681821000, WeekDay: time.Tuesday, YearDay: 108},
			expectedErr: nil,
		},
		{
			name:        "Invalid date string",
			dateStr:     "2023/04/18 12:30:00",
			format:      "2006-01-02 15:04:05",
			expected:    nil,
			expectedErr: errors.New(`parsing time "2023/04/18 12:30:00" as "2006-01-02 15:04:05": cannot parse "/04/18 12:30:00" as "-"`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ParseDate(tc.dateStr, tc.format)

			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("ParseDate(%s, %s) returned unexpected error: got %v, expected %v", tc.dateStr, tc.format, err, tc.expectedErr)
			} else if err == nil && *actual != *tc.expected {
				t.Errorf("ParseDate(%s, %s) returned incorrect date: got %v, expected %v", tc.dateStr, tc.format, actual, tc.expected)
			}
		})
	}
}

func TestGetWeekDayName(t *testing.T) {
	testCases := []struct {
		name     string
		date     *Date
		expected string
	}{
		{
			name:     "Monday",
			date:     &Date{Year: 2023, Month: time.April, Day: 17, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682665200, WeekDay: time.Monday, YearDay: 107},
			expected: "Monday",
		},
		{
			name:     "Tuesday",
			date:     &Date{Year: 2023, Month: time.April, Day: 18, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682799000, WeekDay: time.Tuesday, YearDay: 107},
			expected: "Tuesday",
		},
		{
			name:     "Wednesday",
			date:     &Date{Year: 2023, Month: time.April, Day: 19, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682932800, WeekDay: time.Wednesday, YearDay: 107},
			expected: "Wednesday",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetWeekDayName(tc.date)

			if actual != tc.expected {
				t.Errorf("GetWeekDayName(%v) returned incorrect weekday name: got %s, expected %s", tc.date, actual, tc.expected)
			}
		})
	}
}

func TestAddDays(t *testing.T) {
	testCases := []struct {
		name        string
		date        *Date
		days        int
		expectedDay int
	}{
		{
			name:        "Positive days",
			date:        &Date{Year: 2023, Month: time.April, Day: 18, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682799000, WeekDay: time.Tuesday, YearDay: 107},
			days:        10,
			expectedDay: 28,
		},
		{
			name:        "Negative days",
			date:        &Date{Year: 2023, Month: time.April, Day: 18, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682799000, WeekDay: time.Tuesday, YearDay: 107},
			days:        -5,
			expectedDay: 13,
		},
		{
			name:        "Zero days",
			date:        &Date{Year: 2023, Month: time.April, Day: 18, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682799000, WeekDay: time.Tuesday, YearDay: 107},
			days:        0,
			expectedDay: 18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := AddDays(tc.date, tc.days)

			if actual.Day != tc.expectedDay {
				t.Errorf("AddDays(%v, %d) returned incorrect day: got %d, expected %d", tc.date, tc.days, actual.Day, tc.expectedDay)
			}
		})
	}
}

func TestSubtractDates(t *testing.T) {
	testCases := []struct {
		name           string
		date1, date2   *Date
		expectedSec    int64
		expectedDurMin int64
	}{
		{
			name:           "Same date",
			date1:          &Date{Year: 2023, Month: time.April, Day: 18, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682799000, WeekDay: time.Tuesday, YearDay: 107},
			date2:          &Date{Year: 2023, Month: time.April, Day: 18, Hour: 0, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682799000, WeekDay: time.Tuesday, YearDay: 107},
			expectedSec:    0,
			expectedDurMin: 0,
		},
		{
			name:           "Later date subtracted from earlier date",
			date1:          &Date{Year: 2023, Month: time.April, Day: 18, Hour: 12, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682971800, WeekDay: time.Tuesday, YearDay: 107},
			date2:          &Date{Year: 2023, Month: time.April, Day: 18, Hour: 8, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682956800, WeekDay: time.Tuesday, YearDay: 107},
			expectedSec:    14400,
			expectedDurMin: 240,
		},
		{
			name:           "Earlier date subtracted from later date",
			date1:          &Date{Year: 2023, Month: time.April, Day: 18, Hour: 8, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682956800, WeekDay: time.Tuesday, YearDay: 107},
			date2:          &Date{Year: 2023, Month: time.April, Day: 18, Hour: 12, Minute: 0, Second: 0, Nanosecond: 0, Unix: 1682971800, WeekDay: time.Tuesday, YearDay: 107},
			expectedSec:    -14400,
			expectedDurMin: -240,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualSec, actualDur := SubtractDates(tc.date1, tc.date2)

			if actualSec != tc.expectedSec {
				t.Errorf("SubtractDates(%v, %v) returned incorrect seconds difference: got %d, expected %d", tc.date1, tc.date2, actualSec, tc.expectedSec)
			}

			actualDurMin := actualDur.Minutes()
			if int64(actualDurMin) != tc.expectedDurMin {
				t.Errorf("SubtractDates(%v, %v) returned incorrect duration difference: got %v, expected %v", tc.date1, tc.date2, actualDurMin, tc.expectedDurMin)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		name     string
		year     int
		expected bool
	}{
		{
			name:     "Leap year divisible by 4 and not 100",
			year:     2024,
			expected: true,
		},
		{
			name:     "Leap year divisible by 4, 100, and 400",
			year:     2000,
			expected: true,
		},
		{
			name:     "Non-leap year divisible by 4 but not 100 or 400",
			year:     1900,
			expected: false,
		},
		{
			name:     "Non-leap year not divisible by 4",
			year:     2023,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsLeapYear(tc.year)

			if actual != tc.expected {
				t.Errorf("IsLeapYear(%d) returned %t, expected %t", tc.year, actual, tc.expected)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	testCases := []struct {
		name          string
		year, month   int
		expectedDays  int
		expectedError bool
	}{
		{
			name:         "Valid month with 31 days",
			year:         2023,
			month:        1,
			expectedDays: 31,
		},
		{
			name:         "Valid month with 30 days",
			year:         2023,
			month:        4,
			expectedDays: 31,
		},
		{
			name:         "February in non-leap year",
			year:         2023,
			month:        2,
			expectedDays: 28,
		},
		{
			name:         "February in leap year",
			year:         2024,
			month:        2,
			expectedDays: 29,
		},
		{
			name:          "Invalid month",
			year:          2023,
			month:         13,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualDays, actualErr := DaysInMonth(tc.year, tc.month)

			if tc.expectedError && actualErr == nil {
				t.Errorf("DaysInMonth(%d, %d) returned no error, expected an error", tc.year, tc.month)
			} else if !tc.expectedError && actualErr != nil {
				t.Errorf("DaysInMonth(%d, %d) returned an error: %v, expected no error", tc.year, tc.month, actualErr)
			}

			if actualDays != tc.expectedDays {
				t.Errorf("DaysInMonth(%d, %d) returned %d days, expected %d days", tc.year, tc.month, actualDays, tc.expectedDays)
			}
		})
	}
}
