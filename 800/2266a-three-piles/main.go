package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for range n {
		var p, a, b, c int
		fmt.Scan(&p, &a, &b, &c)
		fmt.Println(max(p-a, p-b, p-c))
	}
}
