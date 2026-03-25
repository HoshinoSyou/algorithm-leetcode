package main

import (
	"fmt"
	"sort"
)

func main() {
	//res := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	//res := threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2)
	res := threeSumClosest([]int{-84, 92, 26, 19, -7, 9, 42, -51, 8, 30, -100, -13, -38}, 78)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	var res int
	var differentTemp int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		fmt.Println(nums, i, nums[i])
		temp := nums[i]
		if i > 0 && temp == nums[i-1] {
			continue
		}

		if sum := temp + nums[i+1] + nums[i+2]; sum > target {
			different := absolute(sum - target)
			fmt.Println(sum, sum-target, different, differentTemp, res)
			if i == 0 {
				res = sum
				differentTemp = different
			} else {
				if differentTemp > different {
					res = sum
					differentTemp = different
				}
			}
			fmt.Println(res)
			break
		} else if sum == target {
			return sum
		}
		if sum := temp + nums[n-1] + nums[n-2]; sum < target {
			different := absolute(sum - target)
			fmt.Println("222222222222222", sum, sum-target, different, differentTemp, res)
			if i == 0 {
				res = sum
				differentTemp = different
			} else {
				if differentTemp > different {
					res = sum
					differentTemp = different
				}
			}
			continue
		} else if sum == target {
			return sum
		}
		j := i + 1
		k := n - 1
		for j < k {
			fmt.Println(nums, i, nums[i], j, nums[j], k, nums[k])
			sum := nums[i] + nums[j] + nums[k]
			different := absolute(sum - target)
			fmt.Println(sum, different, differentTemp, res)
			// 判断和值与target的差值，与之前循环中最低的差值哪个更低
			if i == 0 && j == i+1 && k == n-1 {
				res = sum
				differentTemp = different
			} else {
				if differentTemp > different {
					res = sum
					differentTemp = different
				}
			}
			// 移动j、k指针
			if sum < target {
				j++
				for j > i+1 && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > target {
				k--
				for k < n-1 && nums[k] == nums[k+1] {
					k--
				}
			} else {
				return sum
			}
		}
	}
	return res
}

func absolute(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}
