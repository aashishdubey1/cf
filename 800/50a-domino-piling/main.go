package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	totalSqaures := m * n
	fittedDomino := totalSqaures / 2
	fmt.Println(fittedDomino)
}
