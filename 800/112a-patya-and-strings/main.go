package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	for i := range str1 {
		if str1[i] < str2[i] {
			fmt.Println(-1)
			return
		} else if str1[i] > str2[i] {
			fmt.Println(1)
			return
		} else {
			continue
		}
	}
	fmt.Println(0)
}
