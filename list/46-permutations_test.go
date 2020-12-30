package main

import (
	"fmt"
	"testing"
)

func TestPermutationsDfs(t *testing.T) {
	s := []int{5, 4, 6, 2}
	r := permuteDfs(s)
	fmt.Println(r)
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
