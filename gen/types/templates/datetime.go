package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Precision int

func (p Precision) String() string {
	switch p {
	case DatePrecision:
		return "date"
	case YearMonthPrecision:
		return "year-month"
	case YearPrecision:
		return "year"
	case TimestampPrecision:
		return "timestamp"
	}
	return "timestamp"
}

const (
	TimestampPrecision Precision = iota
	DatePrecision
	YearMonthPrecision
	YearPrecision
)

type DateTime struct {
	Time      time.Time
	Precision Precision
}

func NewDateTime(time time.Time, precision Precision) *DateTime {
	return &DateTime{Time: time, Precision: precision}
}

func Date(year, month, day int) DateTime {
	return DateTime{Time: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), Precision: DatePrecision}
}

func NewDate(year, month, day int) *DateTime {
	d := Date(year, month, day)
	return &d
}

func YearMonth(year, month int) DateTime {
	return DateTime{Time: time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC), Precision: YearMonthPrecision}
}

func NewYearMonth(year, month int) *DateTime {
	d := YearMonth(year, month)
	return &d
}

func Year(year int) DateTime {
	return DateTime{Time: time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC), Precision: YearPrecision}
}

func NewYear(year int) *DateTime {
	d := Year(year)
	return &d
}

func (d DateTime) String() string {
	switch d.Precision {
	case TimestampPrecision:
		return d.Time.Format(time.RFC3339)
	case YearMonthPrecision:
		return d.Time.Format("2006-01")
	case YearPrecision:
		return d.Time.Format("2006")
	case DatePrecision:
		return d.Time.Format("2006-01-02")
	default:
		return ""
	}
}

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	// strData := string(data)
	// if len(data) <= 12 {
	// 	d.Precision = DatePrecision
	// 	d.Time, err = time.ParseInLocation("\"2006-01-02\"", strData, time.Local)
	// 	if err != nil {
	// 		d.Precision = YearMonthPrecision
	// 		d.Time, err = time.ParseInLocation("\"2006-01\"", strData, time.Local)
	// 	}
	// 	if err != nil {
	// 		d.Precision = YearPrecision
	// 		d.Time, err = time.ParseInLocation("\"2006\"", strData, time.Local)
	// 	}
	// 	if err != nil {
	// 		err = fmt.Errorf("unable to parse DateTime: %s", strData)
	// 		d.Precision = TimestampPrecision // DateTimeUnknown
	// 	}
	// } else {
	// 	d.Precision = TimestampPrecision
	// 	d.Time = time.Time{}
	// 	err = d.Time.UnmarshalJSON(data)
	// }

	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	*d, err = ParseDateTime(string(data))
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func ParseDateTime(s string) (d DateTime, err error) {
	s = strings.Trim(s, `"'`)
	if len(s) <= 10 {
		d.Precision = DatePrecision
		d.Time, err = time.ParseInLocation("2006-01-02", s, time.Local)
		if err != nil {
			d.Precision = YearMonthPrecision
			d.Time, err = time.ParseInLocation("2006-01", s, time.Local)
		}
		if err != nil {
			d.Precision = YearPrecision
			d.Time, err = time.ParseInLocation("2006", s, time.Local)
		}
		if err != nil {
			err = fmt.Errorf("unable to parse DateTime: %s", s)
			d.Precision = TimestampPrecision // DateTimeUnknown
		}
	} else {
		d.Precision = TimestampPrecision
		d.Time, err = time.Parse(time.RFC3339, s)
	}
	return d, err
}

// Time.
type Time struct {
	time time.Time
}

func NewTime(hour, minute, second int) *Time {
	return &Time{time: time.Date(0, 1, 1, hour, minute, second, 0, time.UTC)}
}

func (t *Time) Set(hour, minute, second int) {
	t.time = time.Date(0, 1, 1, hour, minute, second, 0, time.UTC)
}

func (t Time) Hour() int {
	return t.time.Hour()
}

func (t Time) Minute() int {
	return t.time.Minute()
}

func (t Time) Second() int {
	return t.time.Second()
}

func (t Time) String() string {
	return t.time.Format("15:04:05")
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	*t, err = ParseTime(string(data))
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func ParseTime(s string) (t Time, err error) {
	time, err := time.ParseInLocation("15:04:05", strings.Trim(s, `"'`), time.UTC)
	if err != nil {
		return t, err
	}
	return Time{time: time}, nil
}
