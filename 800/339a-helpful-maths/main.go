package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	arr := strings.Split(s, "+")
	slices.Sort(arr)
	fs := strings.Join(arr, "+")
	fmt.Println(fs)
}
