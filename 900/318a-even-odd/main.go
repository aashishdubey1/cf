package main

import "fmt"

func main() {

	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	EvenOdd(n, k)

	// O(n) time and O(n) space
	// el := []int{}
	// for i := 1; i <= n; i++ {
	// 	if i%2 != 0 {
	// 		el = append(el, i)
	// 	}
	// }
	// for i := 1; i <= n; i++ {
	// 	if i%2 == 0 {
	// 		el = append(el, i)
	// 	}
	// }
	// fmt.Println(el[k-1])
}

// O(1) itme and space
func EvenOdd(n, k int) {
	oddCount := (n + 1) / 2

	if k > oddCount {
		newK := k - oddCount
		fmt.Println(newK * 2)
	} else {
		fmt.Println(k*2 - 1)
	}
}
