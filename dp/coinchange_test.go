package main

import (
	"math"
	"testing"

	"github.com/soooldier/leetcode/utils"
)

func TestCoinChange(t *testing.T) {
	var cases = []struct {
		input1   []int
		input2   int
		expected int // expescted result
	}{
		{[]int{3, 4, 5, 6, 7}, 1, -1},
		{[]int{3, 4, 5, 6, 7}, 2, -1},
		{[]int{3, 4, 5, 6, 7}, 3, 1},
		{[]int{3, 4, 5, 6, 7}, 8, 2},
		{[]int{3, 4, 5, 6, 7}, 9, 2},
		{[]int{3, 4, 5, 6, 7}, 10, 2},
		{[]int{3, 4, 5, 6, 7}, 14, 2},
	}
	for _, tt := range cases {
		actual1 := coinChangeRecursion(tt.input1, tt.input2)
		if actual1 != tt.expected {
			t.Errorf("coinChangeRecursion(%#v, %#v) = %v; expected %v", tt.input1, tt.input2, actual1, tt.expected)
		}

		actual2 := coinChangeNotes(tt.input1, tt.input2)
		if actual2 != tt.expected {
			t.Errorf("coinChangeNotes(%#v, %#v) = %v; expected %v", tt.input1, tt.input2, actual2, tt.expected)
		}

		actual3 := coinChangeDp(tt.input1, tt.input2)
		if actual3 != tt.expected {
			t.Errorf("coinChangeDp(%#v, %#v) = %v; expected %v", tt.input1, tt.input2, actual3, tt.expected)
		}
	}
}

// 直接递归
func coinChangeRecursion(coins []int, total int) int {
	if total == 0 {
		return 0
	}
	num := math.MaxInt32
	for _, v := range coins {
		if total-v < 0 {
			continue
		}
		sub := coinChangeRecursion(coins, total-v)
		if sub == -1 {
			continue
		}
		num = utils.IntMin(num, sub)
	}
	if num == math.MaxInt32 {
		return -1
	}
	return 1 + num
}

// 备忘录递归
var notes map[int]int

func init() {
	notes = make(map[int]int)
}

func coinChangeNotes(coins []int, total int) int {
	if total == 0 {
		return 0
	}
	if v, ok := notes[total]; ok {
		return v
	}
	num := math.MaxInt32
	for _, v := range coins {
		if total-v < 0 {
			continue
		}
		sub := coinChangeNotes(coins, total-v)
		if sub == -1 {
			continue
		}
		num = utils.IntMin(num, sub)
	}
	if num == math.MaxInt32 {
		notes[total] = -1
		return -1
	}
	notes[total] = 1 + num
	return 1 + num
}

// 动态规划
func coinChangeDp(coins []int, total int) int {
	dp := make([]int, total+1)
	dp[0] = 0
	for i := 1; i <= total; i++ {
		tmp := math.MaxInt32
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			if dp[i-v] == -1 {
				continue
			}
			tmp = utils.IntMin(tmp, dp[i-v])
		}
		if tmp == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = 1 + tmp
		}
	}
	return dp[total]
}
