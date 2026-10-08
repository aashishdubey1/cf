package main

import (
	"fmt"
)

func main() {
	var x int
	var n int
	fmt.Scan(&n)

	for range n {
		var opr string
		fmt.Scan(&opr)
		if opr == "++X" || opr == "X++" {
			x++
		} else {
			x--
		}
	}
	fmt.Println(x)
}
