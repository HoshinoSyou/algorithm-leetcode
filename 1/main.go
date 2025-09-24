package main

import (
	"fmt"
)

func main() {
	res := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(res)
}

// twoSum 哈希表
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			fmt.Println(i, k)
			return []int{i, k}
		}
		m[v] = k
	}
	return nil
}

// twoSum2 两层循环直接解题
func twoSum2(nums []int, target int) []int {
	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			if v+nums[i] != target {
				continue
			}
			return []int{k, i}
		}
	}
	return nil
}
