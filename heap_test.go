package main

import (
	"testing"

	"github.com/soooldier/leetcode/utils"
)

func TestBuildHeap(t *testing.T) {
	var cases = []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 3, 4, 1}, []int{1, 2, 3, 5, 4}},
	}
	for _, tt := range cases {
		actual := insertBuild(tt.input)
		if false == utils.TestIntSliceEq(actual, tt.expected) {
			t.Errorf("insertBuild(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func insertBuild(src []int) []int {
	var dst []int
	dst = append(dst, 0)
	for _, v := range src {
		dst = append(dst, v)
		i := len(dst) - 1
		for i/2 > 0 && dst[i/2] > dst[i] {
			dst[i/2], dst[i] = dst[i], dst[i/2]
			i = i / 2
		}
	}
	return dst[1:]
}
