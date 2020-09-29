package main

import "testing"

func TestValidParentheses(t *testing.T) {
	var cases = []struct {
		input    string // input
		expected bool   // expected result
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"]", false},
		{"", true},
		{"([)]", false},
		{"{[]}", true},
		{"{[]", false},
	}

	for _, tt := range cases {
		actual := isParenthesesValid(tt.input)
		if tt.expected != actual {
			t.Errorf("isParenthesesValid(%s) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

var bmap = map[byte]byte{
	']': '[',
	'}': '{',
	')': '(',
}

func isParenthesesValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var queue []byte
	for i := 0; i < len(s); i++ {
		// 如果第一位是],}或) 则一定不匹配
		if _, exist := bmap[s[i]]; exist && i == 0 {
			return false
		}
		// 左括号入队
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			queue = append(queue, s[i])
		} else if queue[len(queue)-1] == bmap[s[i]] {
			queue = queue[:len(queue)-1]
		}
	}
	return len(queue) == 0
}
