package main

import (
	"testing"

	"github.com/soooldier/leetcode/utils"
)

func TestImplementStrStr(t *testing.T) {
	var cases = []struct {
		input1   string // input
		input2   string // input
		expected int    // expescted result
	}{
		{"mississippi", "issip", 4},
		{"mississippi", "pi", 9},
		{"mississippi", "pii", -1},
		{"pii", "mississippi", -1},
		{"mississippi", "", 0},
		{"", "", -1},
	}
	for _, tt := range cases {
		actual1 := strStrHash(tt.input1, tt.input2)
		if actual1 != tt.expected {
			t.Errorf("strStrHash(%#v, %#v) = %v; expected %v", tt.input1, tt.input2, actual1, tt.expected)
		}
		actual2 := strStrLoop(tt.input1, tt.input2)
		if actual1 != tt.expected {
			t.Errorf("strStrLoop(%#v, %#v) = %v; expected %v", tt.input1, tt.input2, actual2, tt.expected)
		}
	}
}

/*
计算待比较字符串hash值，进而每次遍历比较相同长度字符串的hash值，注意要点：
1. 在遍历计算字符串hash值的时候利用游标移动新增右边字符，剔除左边字符来减少hash计算
2. hash值比较成功的时候需要再挨个比较一下避免hash冲突s
*/
func strStrHash(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n2 > n1 {
		return -1
	}
	hash1 := 0
	hash2 := utils.AsciiHash(needle)
	for i := 0; i < n1; i++ {
		if i == 0 {
			hash1 = utils.AsciiHash(haystack[i:n2])
		} else {
			if n2+i-1 >= n1 {
				return -1
			}
			hash1 = hash1 + int(haystack[n2+i-1]) - int(haystack[i-1])
		}
		if hash1 == hash2 {
			if needle == haystack[i:n2+i] {
				return i
			}
		}
	}
	return -1
}

// 循环遍历比较
func strStrLoop(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n2 > n1 {
		return -1
	}
	if n2 == 0 && n1 > 0 {
		return 0
	}
	for i := 0; i < n1; i++ {
		// 如果needle的字符串长度为3，而haystack字符串已经遍历到最后2个字符，不可能匹配，直接返回-1
		if i >= n1-n2+1 {
			return -1
		}
		for j := 0; j < n2; j++ {
			if needle[j] != haystack[i+j] {
				break
			}
			if j == n2-1 {
				return i
			}
		}
	}
	return -1
}
