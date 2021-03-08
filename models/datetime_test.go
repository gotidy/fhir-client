package models

import (
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {
	tests := []struct {
		name     string
		expected DateTime
		data     string
		wantErr  bool
	}{
		{
			name:     "timestamp",
			expected: DateTime{Time: time.Date(1955, 4, 11, 20, 45, 33, 0, time.Local), Precision: TimestampPrecision},
			data:     `"` + time.Date(1955, 4, 11, 20, 45, 33, 0, time.Local).Format(time.RFC3339) + `"`,
			wantErr:  false,
		},
		{
			name:     "year",
			expected: DateTime{Time: time.Date(1955, 1, 1, 0, 0, 0, 0, time.Local), Precision: YearPrecision},
			data:     `"1955"`,
			wantErr:  false,
		},
		{
			name:     "year-month",
			expected: DateTime{Time: time.Date(1955, 4, 1, 0, 0, 0, 0, time.Local), Precision: YearMonthPrecision},
			data:     `"1955-04"`,
			wantErr:  false,
		},
		{
			name:     "date",
			expected: DateTime{Time: time.Date(1955, 4, 11, 0, 0, 0, 0, time.Local), Precision: DatePrecision},
			data:     `"1955-04-11"`,
			wantErr:  false,
		},
		{
			name:     "not valid",
			expected: DateTime{Time: time.Time{}, Precision: TimestampPrecision},
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
		expected *Time
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
