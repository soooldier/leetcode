package main

import (
	"testing"

	"github.com/soooldier/leetcode/utils"
)

/*
假设背包容量为W，该问题可描述为：
N个物品，在背包允许的容量内，能装下的最大价值物品是多少【每个物品均只有一个】
wt[i],val[i]分别表示第i个物品的重量和价值，i的取值范围是[1...N]；dp[i, w] 表示当背包容量为w时，
放置前i个物品所能获取的最大价值
dp方程：
	dp[i, w] = max(dp[i-1, w], dp[i-1, w-wt[i]] + val[i])
base条件：
	dp[0, ...] = 0【不选择任何物品】
	dp[..., 0] = 0【背包容量为0】
*/

func TestKnapsack(t *testing.T) {
	var cases = []struct {
		input1   []int
		input2   []int
		input3   int
		expected int // expescted result
	}{
		{[]int{0, 5, 5, 3, 4, 3}, []int{0, 40, 50, 20, 30, 35}, 10, 90},
	}
	for _, tt := range cases {
		actual := knapsackDp(tt.input1, tt.input2, tt.input3)
		if actual != tt.expected {
			t.Errorf("knapsackDp(%#v, %#v, %#v) = %v; expected %v", tt.input1, tt.input2, tt.input3, actual, tt.expected)
		}
	}
}

func knapsackDp(w, v []int, cap int) int {
	size := len(w)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, cap+1)
	}
	for i := 0; i < size; i++ {
		for j := 0; j <= cap; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				if j-w[i] < 0 {
					continue
				}
				dp[i][j] = utils.IntMax(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	return dp[size-1][cap]
}
