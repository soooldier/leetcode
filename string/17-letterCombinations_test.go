package main

import (
	"testing"
)

var letters = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func TestLetterCombinations(t *testing.T) {
	var cases = []struct {
		input    string   // input
		expected []string // expected result
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", nil},
		{"02", nil},
		{"20", nil},
	}
	for _, tt := range cases {
		actual := letterCombinationsA(tt.input)
		if false == testEq(actual, tt.expected) {
			t.Errorf("letterCombinationsA(%s) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func letterCombinationsA(digits string) []string {
	var ans []string
	var dfs func(int, string)
	n := len(digits)
	if n == 0 {
		return ans
	}
	dfs = func(step int, part string) {
		if step >= n {
			ans = append(ans, part)
			return
		}
		for _, v := range letters[string(digits[step])] {
			dfs(step+1, part+string(v))
		}
	}
	dfs(0, "")
	return ans
}

func testEq(a, b []string) bool {
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
