package main

import (
	"testing"
)

func TestCoinChange2(t *testing.T) {
	var cases = []struct {
		input1   []int
		input2   int
		expected int // expescted result
	}{
		{[]int{1, 2, 5}, 5, 4},
		{[]int{2}, 3, 0},
		{[]int{10}, 10, 1},
	}
	for _, tt := range cases {
		actual := coinChange2Dp(tt.input1, tt.input2)
		if actual != tt.expected {
			t.Errorf("coinChange2Dp(%#v, %#v) = %v; expected %v", tt.input1, tt.input2, actual, tt.expected)
		}
	}
}

// 动态规划
func coinChange2Dp(coins []int, total int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, total+1)
	}
	for i := 0; i <= len(coins); i++ {
		for j := 0; j <= total; j++ {
			if j == 0 {
				dp[i][0] = 1
			} else if i == 0 {
				dp[0][j] = 0
			} else {
				if j-coins[i-1] < 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
				}
			}
		}
	}
	return dp[len(coins)][total]
}
