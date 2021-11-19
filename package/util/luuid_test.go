package util

import (
	"testing"
	"time"
)

func TestGenerateLifetimeNumber(t *testing.T) {
	type args struct {
		firstName  string
		middleName string
		lastName   string
		gender     string
		dob        time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateLifetimeNumber(tt.args.firstName, tt.args.middleName, tt.args.lastName, tt.args.gender, tt.args.dob); got != tt.want {
				t.Errorf("GenerateLifetimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
