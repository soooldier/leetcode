package main

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var cases = []struct {
		input    string
		expected bool // expescted result
	}{
		{"A maama", true},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
	}
	for _, tt := range cases {
		actual := isPalindrome(tt.input)
		if actual != tt.expected {
			t.Errorf("isPalindrome(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func isPalindrome(s string) bool {
	var s1 string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			s1 += string(s[i])
		}
	}
	s1 = strings.ToLower(s1)
	left, right := 0, len(s1)-1
	for left < right {
		if s1[left] != s1[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
