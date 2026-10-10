package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	textArr := strings.Split(s, "WUB")
	newText := []string{}

	for _, v := range textArr {
		if v != "" {
			newText = append(newText, v)
		}
	}

	el := strings.Join(newText, " ")
	fmt.Println(el)
}
