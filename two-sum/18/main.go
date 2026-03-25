package main

import (
	"fmt"
	"sort"
)

func main() {
	//res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	//res := fourSum([]int{2, 2, 2, 2, 2}, 8)
	res := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		temp := nums[i]
		//fmt.Println(nums, i, temp)
		if i > 0 && temp == nums[i-1] {
			continue
		}
		if temp+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if temp+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			temp2 := nums[j]
			//fmt.Println(nums, i, temp, j, temp2)
			if j > i+1 && temp2 == nums[j-1] {
				continue
			}
			if temp+temp2+nums[j+1]+nums[j+2] > target {
				break
			}
			if temp+temp2+nums[n-1]+nums[n-2] < target {
				continue
			}
			k := j + 1
			m := n - 1
			for k < m {
				sum := temp + temp2 + nums[k] + nums[m]
				//fmt.Println(nums, i, temp, j, temp2, k, nums[k], m, nums[m])
				if sum > target {
					m--
				} else if sum < target {
					k++
				} else {
					res = append(res, []int{temp, temp2, nums[k], nums[m]})
					k++
					//fmt.Println("111111111111111111111111")
					for k < m && k > j+1 && nums[k] == nums[k-1] {
						//fmt.Println("22222222222222222222")
						k++
					}
					m--
					//fmt.Println("3333333333333333")
					for k < m && m < n-1 && nums[m] == nums[m+1] {
						m--
					}
				}
				//fmt.Println("j:", j)
				//fmt.Println("m:", m)
				//fmt.Println(res)
			}
		}
	}
	return res
}
