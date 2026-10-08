package main

import "fmt"

// Link - https://codeforces.com/problemset/problem/71/A

func main() {
	var n int
	fmt.Scan(&n)

	for range n {
		var word string
		fmt.Scan(&word)
		result := wordAbbr(word)
		fmt.Println(result)
	}
}

func wordAbbr(s string) string {
	if len(s) <= 10 {
		return s
	}
	return string(s[0]) + fmt.Sprint(len(s)-2) + string(s[len(s)-1])
}
