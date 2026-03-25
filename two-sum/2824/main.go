package main

import (
	"fmt"
	"sort"
)

func main() {
	//res := countPairs([]int{-1, 1, 2, 3, 1}, 2)
	res := countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2)
	fmt.Println(res)
}

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	res := 0
	for i < j {
		if nums[0]+nums[1] >= target {
			break
		}
		if nums[i]+nums[j] < target {
			res = res + j - i
			i++
		} else {
			j--
		}
	}
	return res
}
