package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Precision int

func (p Precision) String() string {
	switch p {
	case Date:
		return "date"
	case YearMonth:
		return "year-month"
	case Year:
		return "year"
	case Timestamp:
		return "timestamp"
	}
	return "timestamp"
}

const (
	Timestamp Precision = iota
	Date
	YearMonth
	Year
)

type DateTime struct {
	Time      time.Time
	Precision Precision
}

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	strData := string(data)
	if len(data) <= 12 {
		d.Precision = Date
		d.Time, err = time.ParseInLocation("\"2006-01-02\"", strData, time.Local)
		if err != nil {
			d.Precision = YearMonth
			d.Time, err = time.ParseInLocation("\"2006-01\"", strData, time.Local)
		}
		if err != nil {
			d.Precision = Year
			d.Time, err = time.ParseInLocation("\"2006\"", strData, time.Local)
		}
		if err != nil {
			err = fmt.Errorf("unable to parse DateTime: %s", strData)
			d.Precision = Timestamp // DateTimeUnknown
		}
	} else {
		d.Precision = Timestamp
		d.Time = time.Time{}
		err = d.Time.UnmarshalJSON(data)
	}
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	switch d.Precision {
	case Timestamp:
		return json.Marshal(d.Time.Format(time.RFC3339))
	case YearMonth:
		return json.Marshal(d.Time.Format("2006-01"))
	case Year:
		return json.Marshal(d.Time.Format("2006"))
	case Date:
		return json.Marshal(d.Time.Format("2006-01-02"))
	default:
		return nil, fmt.Errorf("unrecognized precision: %s", d.Precision)
	}
}

type Time struct {
	time time.Time
}

func NewTime(hour, minute, second int) Time {
	return Time{time: time.Date(0, 1, 1, hour, minute, second, 0, time.UTC)}
}

func (t *Time) Set(hour, minute, second int) {
	*t = NewTime(hour, minute, second)
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
	time, err := ParseTime(string(data))
	if err != nil {
		return err
	}
	*t = time
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func ParseTime(s string) (t Time, err error) {
	time, err := time.ParseInLocation("\"15:04:05\"", s, time.UTC)
	if err != nil {
		return t, err
	}
	return Time{time: time}, nil
}
