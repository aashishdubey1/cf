package main

import "fmt"

func main() {
	var n int
	var solvable int

	fmt.Scan(&n)

	for range n {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a+b+c >= 2 {
			solvable++
		}
	}
	fmt.Println(solvable)
}
