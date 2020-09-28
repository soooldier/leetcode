package main

import "testing"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func TestLengthOfLongestSubstring(t *testing.T) {
	var cases = []struct {
		input    string // input
		expected int    // expected result
	}{
		{"bbtablud", 6},
		{"abcabcbb", 3},
		{"abba", 2},
		{"pwwkew", 3},
		{"aaaa", 1},
		{"", 0},
	}
	for _, tt := range cases {
		actual := lengthOfLongestSubstringA(tt.input)
		if actual != tt.expected {
			t.Errorf("lengthOfLongestSubstring(%s) = %d; expected %d", tt.input, actual, tt.expected)
		}
	}
}

// 滑动窗口
func lengthOfLongestSubstringA(s string) int {
	sets := make(map[rune]int)
	ans, start := 0, 0
	for k, v := range s {
		if offset, ok := sets[v]; ok && offset >= start {
			start = offset + 1
		}
		sets[v] = k
		ans = max(ans, k-start+1)
	}
	return ans
}

// 滑动窗口（思路更清晰）
func lengthOfLongestSubstringB(s string) int {
	m := make(map[byte]bool)
	n := len(s)
	p, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for p+1 < n && m[s[p+1]] == false {
			m[s[p+1]] = true
			p++
		}
		ans = max(ans, p-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
