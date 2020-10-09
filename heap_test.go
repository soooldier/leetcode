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
		{[]int{1, 9, 2, 8, 3, 7, 4, 6, 5}, []int{1, 3, 2, 5, 8, 7, 4, 9, 6}},
	}
	for _, tt := range cases {
		actual := insertBuild(tt.input)
		if false == utils.TestIntSliceEq(actual, tt.expected) {
			t.Errorf("insertBuild(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}

		actual1 := replaceBuild(tt.input)
		if false == utils.TestIntSliceEq(actual1, tt.expected) {
			t.Errorf("replaceBuild(%#v) = %v; expected %v", tt.input, actual1, tt.expected)
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

func replaceBuild(src []int) []int {
	var dst []int
	dst = append(dst, 0)
	dst = append(dst, src...)
	var heapify func([]int, int)
	heapify = func(s []int, i int) {
		for i/2 > 0 && s[i/2] > s[i] {
			s[i/2], s[i] = s[i], s[i/2]
			i = i / 2
		}
	}
	for i := 1; i < len(dst); i++ {
		heapify(dst, i)
	}
	return dst[1:]
}
