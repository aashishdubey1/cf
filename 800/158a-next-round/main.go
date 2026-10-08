package main

import "fmt"

func main() {
	var n, k, count int
	fmt.Scan(&n)
	fmt.Scan(&k)
	var scores []int

	for range n {
		var val int
		fmt.Scan(&val)
		scores = append(scores, val)
	}

	for _, v := range scores {
		if v >= scores[k-1] && v > 0 {
			count++
		}
	}
	fmt.Println(count)
}
