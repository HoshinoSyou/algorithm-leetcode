package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-100, -70, -60, 110, 120, 130, 160}
	res := threeSum(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		temp := nums[i]
		if i > 0 && temp == nums[i-1] {
			continue
		}
		if temp+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if temp+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := temp + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				result = append(result, []int{temp, nums[j], nums[k]})

				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return result
}
