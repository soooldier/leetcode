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
		{[]int{5, 5, 3, 4, 3}, []int{40, 50, 20, 30, 35}, 10, 90},
	}
	for _, tt := range cases {
		actual := knapsackDp(tt.input1, tt.input2, tt.input3)
		if actual != tt.expected {
			t.Errorf("knapsackDp(%#v, %#v, %#v) = %v; expected %v", tt.input1, tt.input2, tt.input3, actual, tt.expected)
		}
		actual1 := knapsackDpPlus(tt.input1, tt.input2, tt.input3)
		if actual1 != tt.expected {
			t.Errorf("knapsackDpPlus(%#v, %#v, %#v) = %v; expected %v", tt.input1, tt.input2, tt.input3, actual1, tt.expected)
		}
	}
}

// 普通dp
func knapsackDp(w, v []int, cap int) int {
	size := len(w)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, cap+1)
	}
	for i := 0; i < size; i++ {
		for j := 0; j <= cap; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else {
				if i == 0 {
					if w[i] <= j {
						dp[i][j] = v[i]
					} else {
						dp[i][j] = 0
					}
				} else {
					if j-w[i] < 0 {
						dp[i][j] = dp[i-1][j]
					} else {
						dp[i][j] = utils.IntMax(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
					}
				}
			}
		}
	}
	return dp[size-1][cap]
}

// 压缩dp数组版本，空间复杂度更优
func knapsackDpPlus(w, v []int, cap int) int {
	// pre db table
	pdp := make([]int, cap+1)
	// 只放第一个物品
	for i := 0; i <= cap; i++ {
		if w[0] <= i {
			pdp[i] = v[0]
		} else {
			pdp[i] = 0
		}
	}
	var dp []int
	for i := 1; i < len(w); i++ {
		dp = make([]int, cap+1)
		for j := 0; j <= cap; j++ {
			if j-w[i] < 0 {
				dp[j] = pdp[j]
			} else {
				dp[j] = utils.IntMax(pdp[j], pdp[j-w[i]]+v[i])
			}
		}
		pdp = dp
	}
	return pdp[cap]
}
