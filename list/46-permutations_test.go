package main

import (
	"fmt"
	"testing"

	"github.com/soooldier/leetcode/utils"
)

func TestPermutationsDfs(t *testing.T) {
	s := []int{5, 4, 6, 2}
	r := permuteBfs(s)
	fmt.Println(r)
	r1 := permuteDfs(s)
	fmt.Println(r1)
}

func permuteDfs(nums []int) [][]int {
	var list [][]int
	dfs([]int{}, nums, &list)
	return list
}

func dfs(path []int, left []int, list *[][]int) {
	if len(left) == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*list = append(*list, tmp)
		return
	}
	for k, v := range left {
		path = append(path, v)
		left = append(left[:k], left[k+1:]...)
		dfs(path, left, list)
		path = path[:len(path)-1]
		left = append(left[:k], append([]int{v}, left[k:]...)...)
	}
}

func permuteBfs(nums []int) [][]int {
	var list [][]int
	var stack [][]int
	for _, v := range nums {
		stack = append(stack, []int{v})
	}
	bfs(nums, stack, 1, &list)
	return list
}

func bfs(nums []int, stack [][]int, level int, list *[][]int) {
	if level == len(nums) {
		*list = stack
		return
	}
	level++
	var stack1 [][]int
	for _, v := range stack {
		for _, k := range nums {
			if utils.TestInIntSlice(v, k) == false {
				s := make([]int, len(v))
				copy(s, v)
				s = append(s, k)
				stack1 = append(stack1, s)
			}
		}
	}
	bfs(nums, stack1, level, list)
}
