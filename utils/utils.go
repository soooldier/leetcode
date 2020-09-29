package utils

import "sort"

func TestStringSliceEq(a, b []string) bool {
	sort.Sort(sort.StringSlice(a))
	sort.Sort(sort.StringSlice(b))
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
