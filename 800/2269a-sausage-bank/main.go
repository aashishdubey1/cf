package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	for range t {
		var n, k int
		fmt.Scan(&n)
		fmt.Scan(&k)

		total := int(math.Pow(2, float64(n-k+1))) + 2*(k-1)
		fmt.Println(total)
	}
}
