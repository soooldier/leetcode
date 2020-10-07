package main

import (
	"math"
	"testing"
)

func TestMinimumWindowSubstring(t *testing.T) {
	var cases = []struct {
		input    string
		find     string
		expected string // expescted result
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ADOBECODEBANC", "AABC", "ADOBECODEBA"},
		{"ADOBECODEBANC", "", ""},
		{"ADOBECODEBANC", "XYZ", ""},
	}
	for _, tt := range cases {
		actual := minWindow(tt.input, tt.find)
		if actual != tt.expected {
			t.Errorf("minWindow(%#v, %#v) = %v; expected %v", tt.input, tt.find, actual, tt.expected)
		}
	}
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	win, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	start, end := 0, 0
	match, min := 0, math.MaxInt64
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		// 在需要的字符集里面，添加到窗口字符集里面
		if need[c] != 0 {
			win[c]++
			// 如果当前字符的数量匹配需要的字符的数量，则match值+1
			if win[c] == need[c] {
				match++
			}
		}
		// 当所有字符数量都匹配之后，开始缩紧窗口
		for match == len(need) {
			if right-left < min {
				min, start, end = right-left, left, right
			}
			c = s[left]
			left++
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
