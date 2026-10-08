package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)

	distinctChar := make(map[rune]struct{})

	for _, v := range name {
		distinctChar[v] = struct{}{}
	}
	if len(distinctChar)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
