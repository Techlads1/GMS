package util

import (
	"reflect"
	"testing"
)

func TestSortString(t *testing.T) {

	tests := []struct {
		name string
		data []string
		want []string
	}{
		{"unsorted", []string{"exam/login", "exam/store", "exam/index"}, []string{"exam/index", "exam/login", "exam/store"}},
		//{"unsorted", []string{"A", "C", "B"}, []string{"A", "B", "C"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortString(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortString() got = %v, want %v", got, tt.want)
			}

		})
	}
}
