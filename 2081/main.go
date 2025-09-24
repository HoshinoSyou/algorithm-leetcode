package main

import "strconv"

func main() {

}

func kMirror(k int, n int) int64 {
	i := 1
	count := 0
	result := 0
	for {

	}
}

// kMirror2 循环枚举 超时
func kMirror2(k int, n int) int64 {
	i := 1
	count := 0
	result := 0
	for {
		str := strconv.Itoa(i)
		length := len(str)
		strb := []byte(str)
		for k, v := range strb {
			if v != strb[length-k-1] {
				break
			}
			num := strconv.FormatInt(int64(i), k)
			length2 := len(num)
			str2b := []byte(num)
			for k, v := range str2b {
				if v != str2b[length2-k-1] {
					break
				}
			}
		}
		result += i
		if count == n {
			break
		}
		i++
	}
	return int64(result)
}
