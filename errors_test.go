package fhir

import (
	"fmt"
	"testing"
)

func TestIsNotFound(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Not NotFound error",
			args: args{err: fmt.Errorf("some error")},
			want: false,
		},
		{
			name: "NotFound error",
			args: args{err: NewNotFoundError("cea60c25-0d97-4e63-87ec-aae7bcb06785", "Patient")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotFoundError(tt.args.err); got != tt.want {
				t.Errorf("IsNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}
