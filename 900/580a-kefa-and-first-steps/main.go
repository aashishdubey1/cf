package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)

	for i := range n {
		fmt.Scan(&nums[i])
	}

	maxCount := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > maxCount {
			maxCount = count
		}
	}

	fmt.Println(maxCount)
}
