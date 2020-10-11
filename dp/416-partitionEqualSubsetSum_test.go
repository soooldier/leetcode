package main

import "testing"

func TestPartitionEqualSubsetSum(t *testing.T) {
	var cases = []struct {
		input    []int
		expected bool // expescted result
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, tt := range cases {
		actual := canPartition(tt.input)
		if actual != tt.expected {
			t.Errorf("canPartition(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	size := len(nums)
	dp := make([][]bool, size+1)
	for i := 0; i <= size; i++ {
		dp[i] = make([]bool, sum/2+1)
	}
	for i := 0; i <= size; i++ {
		for j := 0; j <= sum/2; j++ {
			if j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[size][sum/2]
}
