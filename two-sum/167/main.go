package main

func main() {

}

// twoSum 双指针
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		}
		if numbers[i]+numbers[j] < target {
			i++
		}
		if numbers[i]+numbers[j] > target {
			j--
		}
	}
	return []int{i + 1, j + 1}
}

// twoSum2 两层循环直接解题
func twoSum2(numbers []int, target int) []int {
	for k, v := range numbers {
		for i := k + 1; i < len(numbers); i++ {
			if v+numbers[i] != target {
				continue
			}
			return []int{k + 1, i + 1}
		}
	}
	return nil
}
