package models

import (
	"testing"
	"time"
)

// func TestFHIRDateTime( *check.C) {
// 	simple := &Simple{}

// 	data := []byte("{ \"foo\": [\"1991-02-01T10:00:00-05:00\", \"1992-02-01\", \"1993-02-01T10:00:00-05:00\", \"1992-06\", \"03:04:05\", \"1994\"]}")
// 	err := json.Unmarshal(data, &simple)
// 	util.CheckErr(err)

// 	c.Assert(simple.Foo, check.HasLen, 6)
// 	loc, err := time.LoadLocation("America/New_York")

// 	c.Assert(simple.Foo[0].Time.Equal(time.Date(1991, time.February, 1, 10, 0, 0, 0, loc)), check.Equals, true)
// 	c.Assert(simple.Foo[0].Precision, check.Equals, Precision(Timestamp))

// 	c.Assert(simple.Foo[1].Time.Equal(time.Date(1992, time.February, 1, 0, 0, 0, 0, time.Local)), check.Equals, true)
// 	c.Assert(simple.Foo[1].Precision, check.Equals, Precision(Date))

// 	c.Assert(simple.Foo[2].Time.Equal(time.Date(1993, time.February, 1, 10, 0, 0, 0, loc)), check.Equals, true)
// 	c.Assert(simple.Foo[2].Precision, check.Equals, Precision(Timestamp))

// 	c.Assert(simple.Foo[3].Time.Equal(time.Date(1992, time.June, 1, 0, 0, 0, 0, time.Local)), check.Equals, true)
// 	c.Assert(simple.Foo[3].Precision, check.Equals, Precision(YearMonth))

// 	c.Assert(simple.Foo[4].Time.Equal(time.Date(0, time.January, 1, 03, 04, 05, 0, time.Local)), check.Equals, true)
// 	c.Assert(simple.Foo[4].Precision, check.Equals, Precision(Time))

// 	c.Assert(simple.Foo[5].Time.Equal(time.Date(1994, time.January, 1, 0, 0, 0, 0, time.Local)), check.Equals, true)
// 	c.Assert(simple.Foo[5].Precision, check.Equals, Precision(Year))

// 	foo0, err := json.Marshal(simple.Foo[0])
// 	c.Assert(string(foo0), check.Equals, "\"1991-02-01T10:00:00-05:00\"")

// 	foo1, err := json.Marshal(simple.Foo[1])
// 	c.Assert(string(foo1), check.Equals, "\"1992-02-01\"")

// 	foo2, err := json.Marshal(simple.Foo[2])
// 	c.Assert(string(foo2), check.Equals, "\"1993-02-01T10:00:00-05:00\"")

// 	foo3, err := json.Marshal(simple.Foo[3])
// 	c.Assert(string(foo3), check.Equals, "\"1992-06\"")

// 	foo4, err := json.Marshal(simple.Foo[4])
// 	c.Assert(string(foo4), check.Equals, "\"03:04:05\"")

// 	foo5, err := json.Marshal(simple.Foo[5])
// 	c.Assert(string(foo5), check.Equals, "\"1994\"")

// 	// TODO: test error handling
// }

func TestDateTime(t *testing.T) {
	tests := []struct {
		name     string
		expected DateTime
		data     string
		wantErr  bool
	}{
		{
			name:     "timestamp",
			expected: DateTime{Time: time.Date(1955, 4, 11, 20, 45, 33, 0, time.Local), Precision: Timestamp},
			data:     `"` + time.Date(1955, 4, 11, 20, 45, 33, 0, time.Local).Format(time.RFC3339) + `"`,
			wantErr:  false,
		},
		{
			name:     "year",
			expected: DateTime{Time: time.Date(1955, 1, 1, 0, 0, 0, 0, time.Local), Precision: Year},
			data:     `"1955"`,
			wantErr:  false,
		},
		{
			name:     "year-month",
			expected: DateTime{Time: time.Date(1955, 4, 1, 0, 0, 0, 0, time.Local), Precision: YearMonth},
			data:     `"1955-04"`,
			wantErr:  false,
		},
		{
			name:     "date",
			expected: DateTime{Time: time.Date(1955, 4, 11, 0, 0, 0, 0, time.Local), Precision: Date},
			data:     `"1955-04-11"`,
			wantErr:  false,
		},
		{
			name:     "not valid",
			expected: DateTime{Time: time.Time{}, Precision: Timestamp},
			data:     `"XXXX"`,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DateTime{}
			if err := d.UnmarshalJSON([]byte(tt.data)); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if d.Precision != tt.expected.Precision {
				t.Errorf("expected Precision: %v; actual: %v", tt.expected.Precision, d.Precision)
				return
			}
			if d.Time != tt.expected.Time {
				t.Errorf("expected Time: %v; actual: %v", tt.expected.Time, d.Time)
				return
			}
			if tt.wantErr {
				return
			}
			data, err := d.MarshalJSON()
			if err != nil {
				t.Errorf("DateTime.MarshalJSON() error = %v", err)
				return
			}
			if string(data) != tt.data {
				t.Errorf("expected: %v; actual: %v", tt.data, string(data))
			}
		})
	}
}

func TestTime(t *testing.T) {
	tests := []struct {
		name     string
		expected Time
		data     string
		wantErr  bool
	}{
		{
			name:     "time",
			expected: NewTime(12, 24, 48),
			data:     `"12:24:48"`,
			wantErr:  false,
		},
		{
			name:     "not valid",
			expected: NewTime(0, 0, 0),
			data:     `"XXXX"`,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Time{}
			if err := d.UnmarshalJSON([]byte(tt.data)); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if d.String() != tt.expected.String() {
				t.Errorf("expected Time: %v; actual: %v", tt.expected, d)
				return
			}
			if tt.wantErr {
				return
			}
			data, err := d.MarshalJSON()
			if err != nil {
				t.Errorf("Time.MarshalJSON() error = %v", err)
				return
			}
			if string(data) != tt.data {
				t.Errorf("expected: %v; actual: %v", tt.data, string(data))
			}
		})
	}
}
