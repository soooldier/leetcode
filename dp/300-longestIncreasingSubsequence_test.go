package main

import (
	"testing"

	"github.com/soooldier/leetcode/utils"
)

/*
https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484498&idx=1&sn=df58ef249c457dd50ea632f7c2e6e761&chksm=9bd7fa5aaca0734c29bcf7979146359f63f521e3060c2acbf57a4992c887aeebe2a9e4bd8a89&scene=178#rd
*/

func TestLongestIncreasingSubsequence(t *testing.T) {
	var cases = []struct {
		input    []int
		expected int // expescted result
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{10, 9, 2}, 1},
		{[]int{}, 0},
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
	}
	for _, tt := range cases {
		actual := lengthOfLIS(tt.input)
		if actual != tt.expected {
			t.Errorf("lengthOfLIS(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func lengthOfLIS(nums []int) int {
	ans := 0
	if len(nums) == 0 {
		return ans
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 先初始化为1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = utils.IntMax(dp[i], dp[j]+1)
			}
		}
		ans = utils.IntMax(ans, dp[i])
	}
	return ans
}
