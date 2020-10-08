package main

import "testing"

func TestReorganizeString(t *testing.T) {
	var cases = []struct {
		input    string
		expected string // expescted result
	}{
		{"aab", "aba"},
		{"aaab", ""},
		{"", ""},
	}
	for _, tt := range cases {
		actual := reorganizeString(tt.input)
		if actual != tt.expected {
			t.Errorf("reorganizeString(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func reorganizeString(S string) string {
	var ans string
	return ans
}
