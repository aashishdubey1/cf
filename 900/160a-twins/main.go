package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	coins := make([]int, n)
	totalSum := 0

	for i := range n {
		fmt.Scan(&coins[i])
		totalSum += coins[i]
	}

	slices.Sort(coins)
	slices.Reverse(coins)

	var coinsValue int
	remainingSum := totalSum

	for i, v := range coins {
		coinsValue += v
		remainingSum = totalSum - coinsValue

		if coinsValue > remainingSum {
			fmt.Println(i + 1)
			return
		}
	}
}
