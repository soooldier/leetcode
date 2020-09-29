package main

import (
	"testing"

	"github.com/soooldier/leetcode/utils"
)

/*
关键点：
	任意时刻剩余左括号的个数要小于等于右括号个数才能正常生成
*/
func TestGenerateParenthesis(t *testing.T) {
	var cases = []struct {
		input    int      // input
		expected []string // expescted result
	}{
		{0, nil},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range cases {
		actual1 := generateParenthesisDfs(tt.input)
		if false == utils.TestStringSliceEq(actual1, tt.expected) {
			t.Errorf("generateParenthesisDfs(%d) = %v; expected %v", tt.input, actual1, tt.expected)
		}

		actual2 := generateParenthesisBfs(tt.input)
		if false == utils.TestStringSliceEq(actual2, tt.expected) {
			t.Errorf("generateParenthesisBfs(%d) = %v; expected %v", tt.input, actual2, tt.expected)
		}
	}
}

// dfs + 剪枝
func generateParenthesisDfs(n int) []string {
	var ans []string
	if n == 0 {
		return ans
	}
	var dfs func(int, int, string)
	dfs = func(left, right int, s string) {
		if left == 0 && right == 0 {
			ans = append(ans, s)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			dfs(left-1, right, s+"(")
		}
		if right > 0 {
			dfs(left, right-1, s+")")
		}
	}
	dfs(n, n, "")
	return ans
}

// 广度优先搜索
func generateParenthesisBfs(n int) []string {
	if n == 0 {
		return nil
	}
	queue := []string{"()"}
	for i := 1; i < n; i++ {
		size := len(queue)
		var tqueue []string
		for j := 0; j < size; j++ {
			// case 1 先放左括号，随后放队列里的括号对，最后再放右括号
			case1 := "(" + queue[j] + ")"
			// case 2 先放左括号，再放右括号，最后放队列里的括号对
			case2 := "()" + queue[j]
			// case 3 先放队列里的括号对，再放放左括号，最后右括号
			case3 := queue[j] + "()"
			tqueue = append(tqueue, case1)
			if case2 != case3 {
				tqueue = append(tqueue, case2)
				tqueue = append(tqueue, case3)
			} else {
				tqueue = append(tqueue, case2)
			}
		}
		queue = tqueue
	}
	return queue
}
