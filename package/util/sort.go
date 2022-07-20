package util

import (
	"sort"
)

func SortString(data []string) []string {
	sort.Strings(data)
	return data
}
