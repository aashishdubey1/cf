package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	yearCounter := 0

	for {
		if a > b {
			fmt.Println(yearCounter)
			return
		}
		a *= 3
		b *= 2
		yearCounter++
	}
}
