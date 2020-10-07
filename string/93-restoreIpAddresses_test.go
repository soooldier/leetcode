package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/soooldier/leetcode/utils"
)

func TestRestoreIpAddresses(t *testing.T) {
	var cases = []struct {
		input    string
		expected []string // expescted result
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"0000", []string{"0.0.0.0"}},
		{"1111", []string{"1.1.1.1"}},
		{"010010", []string{"0.10.0.10", "0.100.1.0"}},
	}
	for _, tt := range cases {
		actual := restoreIpAddresses(tt.input)
		if false == utils.TestStringSliceEq(actual, tt.expected) {
			t.Errorf("restoreIpAddresses(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func restoreIpAddresses(s string) []string {
	segments := make([]string, 4)
	ans := []string{}
	var dfs func(s string, index int, carry int)
	dfs = func(s string, index int, carry int) {
		if len(s[index:]) > (4-carry)*3 {
			// 剩余字符太多，不可能拼接成功直接剪枝
			return
		}
		if carry == 4 {
			if index == len(s) {
				// 已经找到4段并且字符串s已经用完
				ans = append(ans, strings.Join(segments, "."))
				return
			} else {
				// 找到4段字符后还有剩余字符
				return
			}
		}
		if index == len(s) {
			// 字符串s已经用完了但还没拼够4段字符
			return
		}
		if s[index] == '0' {
			// 如果第一个字符是0，那么这段字符只能是0
			segments[carry] = "0"
			dfs(s, index+1, carry+1)
		}
		var addr int
		for i := index; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr < 256 {
				segments[carry] = strconv.Itoa(addr)
				dfs(s, i+1, carry+1)
			} else {
				break
			}
		}
		return
	}
	dfs(s, 0, 0)
	return ans
}
