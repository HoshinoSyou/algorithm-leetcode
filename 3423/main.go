package main

import "fmt"

func main() {
	//res := maxAdjacentDistance([]int{1, 2, 4})
	res := maxAdjacentDistance([]int{-5, -10, -5})
	fmt.Println(res)
}

func maxAdjacentDistance(nums []int) int {
	var result int
	for k, v := range nums {
		if k == 0 {
			k = len(nums)
		}
		temp := real(v - nums[k-1])
		if result < temp {
			result = temp
		}
	}
	return result
}

func real(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
