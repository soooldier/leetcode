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

func TestIntSliceEq(a, b []int) bool {
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

func AsciiHash(s string) int {
	var i int
	for _, v := range s {
		i = i + int(v)
	}
	return i
}

func IntMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func IntMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}
